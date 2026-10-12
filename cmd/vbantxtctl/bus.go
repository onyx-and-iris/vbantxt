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
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Bus[idx].Mute()
		fmt.Fprintf(ctx.Out, "bus %d mute: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("bus[%d].mute=%d", idx, val)); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "bus %d mute set: %t\n", idx, *cmd.State == "true")
	return nil
}
