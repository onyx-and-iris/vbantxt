package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// StripEqCmdGroup defines the command group for controlling the EQ settings of a specific strip.
type StripEqCmdGroup struct {
	On      StripEqOnCmd      `cmd:"" help:"Get or set strip EQ on/off."`
	Channel StripEqChannelCmd `cmd:"" help:"Strip EQ channel operations."`
}

// StripEqOnCmd defines the command for getting or setting the EQ on/off state of a strip.
type StripEqOnCmd struct {
	State *string `arg:"" optional:"" enum:"true,false" help:"Set EQ on/off. If omitted, show current state."`
}

// Run executes the StripEqOnCmd command, either retrieving the current EQ on/off state of the strip or setting it based on the provided argument.
func (cmd *StripEqOnCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		on := p.States().Strip[idx].EQOn()

		fmt.Fprintf(ctx.Out, "strip %d eq.on: %t\n", idx, on)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].eq.on=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set eq.on state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d eq.on set: %s\n", idx, *cmd.State)
	return nil
}

// StripEqChannelCmd defines the command group for controlling a specific EQ channel of a strip.
type StripEqChannelCmd struct {
	Index struct {
		Index int                   `arg:"" help:"EQ channel index (0-based)."`
		Cell  StripEqChannelCellCmd `cmd:"" help:"Strip EQ channel cell operations."`
	} `help:"Control a specific EQ channel of the strip by index." arg:""`
}

// StripEqChannelCellCmd defines the command group for controlling a specific EQ channel cell of a strip.
type StripEqChannelCellCmd struct {
	Index struct {
		Index int                `arg:"" help:"EQ cell index (0-based)."`
		On    StripEqCellOnCmd   `cmd:"" help:"Get/set EQ cell on/off."`
		Type  StripEqCellTypeCmd `cmd:"" help:"Get/set EQ cell type."`
		Freq  StripEqCellFreqCmd `cmd:"" help:"Get/set EQ cell frequency."`
		Gain  StripEqCellGainCmd `cmd:"" help:"Get/set EQ cell gain."`
		Q     StripEqCellQCmd    `cmd:"" help:"Get/set EQ cell Q factor."`
	} `help:"Control a specific EQ channel cell of the strip by index." arg:""`
}

// StripEqCellOnCmd defines the command for getting or setting the on/off state of a specific EQ channel cell of a strip.
type StripEqCellOnCmd struct {
	State *string `arg:"" optional:"" enum:"true,false"`
}

// Run executes the StripEqCellOnCmd command, either retrieving the current on/off state of the EQ cell or setting it based on the provided argument.
func (cmd *StripEqCellOnCmd) Run(
	ctx *appContext,
	strip *StripCmdGroup,
	ch *StripEqChannelCmd,
	cell *StripEqChannelCellCmd,
) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		if ch.Index.Index != 0 {
			return fmt.Errorf("only EQ channel 0 is supported")
		}

		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		on := p.Strip[idx].PEQ_On[cell.Index.Index] != 0

		fmt.Fprintf(
			ctx.Out,
			"strip %d eq channel %d cell %d on: %t\n",
			idx,
			ch.Index.Index,
			cell.Index.Index,
			on,
		)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf(
		"strip[%d].eq.channel[%d].cell[%d].on=%d",
		idx, ch.Index.Index, cell.Index.Index, val,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell on/off: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"strip %d eq channel %d cell %d on set: %t\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State == "true",
	)
	return nil
}

// StripEqCellTypeCmd defines the command for getting or setting the type of a specific EQ channel cell of a strip.
type StripEqCellTypeCmd struct{ State *string }

// StripEqCellFreqCmd defines the command for getting or setting the frequency of a specific EQ channel cell of a strip.
type StripEqCellFreqCmd struct{ State *float64 }

// StripEqCellGainCmd defines the command for getting or setting the gain of a specific EQ channel cell of a strip.
type StripEqCellGainCmd struct{ State *float64 }

// StripEqCellQCmd defines the command for getting or setting the Q factor of a specific EQ channel cell of a strip.
type StripEqCellQCmd struct{ State *float64 }
