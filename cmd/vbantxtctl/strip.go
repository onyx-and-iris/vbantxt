package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// StripCmdGroup defines the command group for controlling a specific strip, including its mute state and EQ settings.
type StripCmdGroup struct {
	Index struct {
		Index int          `arg:"" help:"The index of the strip. (0-based indexing)"`
		Mono  StripMonoCmd `cmd:""       help:"Get or set the mono state of the strip."`
		Solo  StripSoloCmd `cmd:""       help:"Get or set the solo state of the strip."`
		Mute  StripMuteCmd `cmd:""       help:"Get or set the mute state of the strip."`

		Eq StripEqCmdGroup `cmd:"" help:"Get or set the EQ settings of the strip."`
	} `arg:"" help:"Control a specific strip by index."`
}

// StripMonoCmd defines the command for getting or setting the mono state of a strip.
type StripMonoCmd struct {
	State *string `arg:"" help:"The mono state to set (true or false). If not provided, the current mono state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripMonoCmd command, either retrieving the current mono state of the strip or setting it based on the provided argument.
func (cmd *StripMonoCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].Mono()
		fmt.Fprintf(ctx.Out, "strip %d mono: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("strip[%d].mono=%d", idx, val)); err != nil {
		return fmt.Errorf("failed to set mono state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "strip %d mono set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripSoloCmd defines the command for getting or setting the solo state of a strip.
type StripSoloCmd struct {
	State *string `arg:"" help:"The solo state to set (true or false). If not provided, the current solo state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripSoloCmd command, either retrieving the current solo state of the strip or setting it based on the provided argument.
func (cmd *StripSoloCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].Solo()
		fmt.Fprintf(ctx.Out, "strip %d solo: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("strip[%d].solo=%d", idx, val)); err != nil {
		return fmt.Errorf("failed to set solo state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "strip %d solo set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripMuteCmd defines the command for getting or setting the mute state of a strip.
type StripMuteCmd struct {
	State *string `arg:"" help:"The mute state to set (true or false). If not provided, the current mute state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripMuteCmd command, either retrieving the current mute state of the strip or setting it based on the provided argument.
func (cmd *StripMuteCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].Mute()
		fmt.Fprintf(ctx.Out, "strip %d mute: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("strip[%d].mute=%d", idx, val)); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "strip %d mute set: %t\n", idx, *cmd.State == "true")
	return nil
}
