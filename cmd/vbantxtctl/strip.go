package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// StripCmdGroup defines the command group for controlling a specific strip, including its mute state and EQ settings.
type StripCmdGroup struct {
	Index struct {
		Index int          `arg:"" help:"The index of the strip. (0-based indexing)"`
		Mute  StripMuteCmd `cmd:""       help:"Get or set the mute state of the strip."`

		Eq StripEqCmdGroup `cmd:"" help:"Get or set the EQ settings of the strip."`
	} `arg:"" help:"Control a specific strip by index."`
}

// StripMuteCmd defines the command for getting or setting the mute state of a strip.
type StripMuteCmd struct {
	State *string `arg:"" help:"The mute state to set (true or false). If not provided, the current mute state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripMuteCmd command, either retrieving the current mute state of the strip or setting it based on the provided argument.
func (cmd *StripMuteCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		if p.States().Strip[strip.Index.Index].Mute() {
			fmt.Fprintf(ctx.Out, "Strip %d is muted\n", strip.Index.Index)
		} else {
			fmt.Fprintf(ctx.Out, "Strip %d is unmuted\n", strip.Index.Index)
		}
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	} else {
		val = 0
	}
	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("strip[%d].mute=%d", strip.Index.Index, val)); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "Strip %d mute set: %s\n", strip.Index.Index, *cmd.State)
	return nil
}
