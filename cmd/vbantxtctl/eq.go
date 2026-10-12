package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

type EqCmdGroup struct {
	On    EqOnCmd `cmd:"" help:"Get or set the EQ on/off state of the strip."`
	Index struct {
		Index   int          `arg:"" help:"The index of the EQ strip. (0-based indexing)"`
		Channel EqChannelCmd `cmd:"" help:"Get or set the EQ channel of the strip."`
	} `       help:"Control a specific EQ strip by index."        arg:""`
}

type EqOnCmd struct {
	State *string `arg:"" help:"The EQ on/off state to set (true or false). If not provided, the current EQ on/off state will be returned." optional:"" enum:"true,false"`
}

func (cmd *EqOnCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		if p.States().Strip[strip.Index.Index].EQOn() {
			fmt.Fprintf(ctx.Out, "Strip %d EQ is on\n", strip.Index.Index)
		} else {
			fmt.Fprintf(ctx.Out, "Strip %d EQ is off\n", strip.Index.Index)
		}
		return nil
	}

	var eqValue int
	if *cmd.State == "true" {
		eqValue = 1
	} else {
		eqValue = 0
	}
	if err := ctx.Client.SendTXT(ctx.Ctx, fmt.Sprintf("strip[%d].eq.on=%d", strip.Index.Index, eqValue)); err != nil {
		return fmt.Errorf("failed to set EQ state: %w", err)
	}
	fmt.Fprintf(ctx.Out, "Strip %d EQ set: %s\n", strip.Index.Index, *cmd.State)
	return nil
}

type EqChannelCmd struct {
	Index struct {
		Index int                   `arg:"" help:"The index of the EQ channel. (0-based indexing)"`
		Cell  EqChannelCellCmdGroup `cmd:"" help:"Get or set the EQ channel of the strip."`
	} `help:"Control a specific EQ channel of the strip by index." arg:""`
}

type EqChannelCellCmdGroup struct {
	Index struct {
		Index int                       `arg:"" help:"The index of the EQ channel cell. (0-based indexing)"`
		On    EqChannelCellOnCmd        `cmd:"" help:"Get or set the EQ channel cell on/off state."`
		Type  EqChannelCellTypeCmd      `cmd:"" help:"Get or set the EQ channel cell type."`
		Freq  EqChannelCellFrequencyCmd `cmd:"" help:"Get or set the EQ channel cell frequency."`
		Gain  EqChannelCellGainCmd      `cmd:"" help:"Get or set the EQ channel cell gain."`
		Q     EqChannelCellQCmd         `cmd:"" help:"Get or set the EQ channel cell Q factor."`
	} `help:"Control a specific EQ channel cell of the strip by index." arg:""`
}

type EqChannelCellOnCmd struct {
	State *string `arg:"" help:"The EQ channel cell on/off state to set (true or false). If not provided, the current state will be returned." optional:"" enum:"true,false"`
}

func (cmd *EqChannelCellOnCmd) Run(
	ctx *appContext,
	strip *StripCmdGroup,
	ch *EqChannelCmd,
	cell *EqChannelCellCmdGroup,
) error {
	if cmd.State == nil {
		if ch.Index.Index != 0 {
			return fmt.Errorf("only EQ channel 0 is supported")
		}

		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		if p.Strip[strip.Index.Index].PEQ_On[cell.Index.Index] != 0 {
			fmt.Fprintf(ctx.Out, "Strip %d EQ is on\n", strip.Index.Index)
		} else {
			fmt.Fprintf(ctx.Out, "Strip %d EQ is off\n", strip.Index.Index)
		}
		return nil
	}
	return nil
}

type EqChannelCellTypeCmd struct {
	State *string `arg:"" help:"The EQ channel cell type to set. If not provided, the current type will be returned." optional:""`
}

type EqChannelCellFrequencyCmd struct {
	State *float64 `arg:"" help:"The EQ channel cell frequency to set. If not provided, the current frequency will be returned." optional:""`
}

type EqChannelCellGainCmd struct {
	State *float64 `arg:"" help:"The EQ channel cell gain to set. If not provided, the current gain will be returned." optional:""`
}

type EqChannelCellQCmd struct {
	State *float64 `arg:"" help:"The EQ channel cell Q factor to set. If not provided, the current Q factor will be returned." optional:""`
}
