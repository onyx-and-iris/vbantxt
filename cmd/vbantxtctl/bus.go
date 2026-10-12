package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// BusCmdGroup defines the command group for controlling a specific bus, including its mute state and EQ settings.
type BusCmdGroup struct {
	Index struct {
		Index int        `arg:"" help:"The index of the bus. (0-based indexing)"`
		Mute  BusMuteCmd `cmd:""       help:"Get or set the mute state of the bus."`

		Eq BusEqCmdGroup `cmd:"" help:"Get or set the EQ settings of the bus."`
	} `arg:"" help:"Control a specific bus by index."`
}

// BusMuteCmd defines the command for getting or setting the mute state of a bus.
type BusMuteCmd struct {
	State *string `arg:"" help:"The mute state to set (true or false). If not provided, the current mute state will be returned." optional:"" enum:"true,false"`
}

// Run executes the BusMuteCmd command, either retrieving the current mute state of the bus or setting it based on the provided argument.
func (cmd *BusMuteCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		if p.States().Bus[bus.Index.Index].Mute() {
			fmt.Fprintf(ctx.Out, "Bus %d is muted\n", bus.Index.Index)
		} else {
			fmt.Fprintf(ctx.Out, "Bus %d is unmuted\n", bus.Index.Index)
		}
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	} else {
		val = 0
	}
	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("bus[%d].mute=%d", bus.Index.Index, val)); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "Bus %d mute set: %s\n", bus.Index.Index, *cmd.State)
	return nil
}
