package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// BusEqCmdGroup defines the command group for controlling the EQ settings of a specific bus.
type BusEqCmdGroup struct {
	On      BusEqOnCmd      `cmd:"" help:"Get or set bus EQ on/off."`
	Channel BusEqChannelCmd `cmd:"" help:"Bus EQ channel operations."`
}

// BusEqOnCmd defines the command for getting or setting the EQ on/off state of a bus.
type BusEqOnCmd struct {
	State *string `arg:"" optional:"" enum:"true,false"`
}

// Run executes the BusEqOnCmd command, either retrieving the current EQ on/off state of the bus or setting it based on the provided argument.
func (cmd *BusEqOnCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		on := p.States().Bus[idx].EQOn()

		fmt.Fprintf(ctx.Out, "bus %d eq.on: %t\n", idx, on)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("bus[%d].eq.on=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set eq.on state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d eq.on set: %t\n", idx, *cmd.State == "true")
	return nil
}

// BusEqChannelCmd defines the command group for controlling a specific EQ channel of a bus.
type BusEqChannelCmd struct {
	Index struct {
		Index int                 `arg:""`
		Cell  BusEqChannelCellCmd `cmd:""`
	} `help:"Control a specific EQ channel of the bus by index." arg:""`
}

// BusEqChannelCellCmd defines the command group for controlling a specific EQ channel cell of a bus.
type BusEqChannelCellCmd struct {
	Index struct {
		Index int              `arg:""`
		On    BusEqCellOnCmd   `cmd:""`
		Type  BusEqCellTypeCmd `cmd:""`
		Freq  BusEqCellFreqCmd `cmd:""`
		Gain  BusEqCellGainCmd `cmd:""`
		Q     BusEqCellQCmd    `cmd:""`
	} `help:"Control a specific EQ channel cell of the bus by index." arg:""`
}

// BusEqCellOnCmd defines the command for getting or setting the on/off state of a specific EQ channel cell of a bus.
type BusEqCellOnCmd struct {
	State *string `arg:"" optional:"" enum:"true,false"`
}

// Run executes the BusEqCellOnCmd command, either retrieving the current on/off state of the EQ cell or setting it based on the provided argument.
func (cmd *BusEqCellOnCmd) Run(
	ctx *appContext,
	bus *BusCmdGroup,
	ch *BusEqChannelCmd,
	cell *BusEqChannelCellCmd,
) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		return fmt.Errorf("bus eq channel cell on queries are not supported")
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf(
		"bus[%d].eq.channel[%d].cell[%d].on=%d",
		idx, ch.Index.Index, cell.Index.Index, val,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell on: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"bus %d eq channel %d cell %d on set: %t\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State == "true",
	)
	return nil
}

// BusEqCellTypeCmd defines the command for getting or setting the type of a specific EQ channel cell of a bus.
type BusEqCellTypeCmd struct {
	State *string
}

func (cmd *BusEqCellTypeCmd) Run(
	ctx *appContext,
	bus *BusCmdGroup,
	ch *BusEqChannelCmd,
	cell *BusEqChannelCellCmd,
) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		return fmt.Errorf("bus eq channel cell type queries are not supported")
	}

	txt := fmt.Sprintf(
		"bus[%d].eq.channel[%d].cell[%d].type=%s",
		idx, ch.Index.Index, cell.Index.Index, *cmd.State,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell type: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"bus %d eq channel %d cell %d type set: %s\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State,
	)
	return nil
}

// BusEqCellFreqCmd defines the command for getting or setting the frequency of a specific EQ channel cell of a bus.
type BusEqCellFreqCmd struct {
	State *float64
}

func (cmd *BusEqCellFreqCmd) Run(
	ctx *appContext,
	bus *BusCmdGroup,
	ch *BusEqChannelCmd,
	cell *BusEqChannelCellCmd,
) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		return fmt.Errorf("bus eq channel cell frequency queries are not supported")
	}

	txt := fmt.Sprintf(
		"bus[%d].eq.channel[%d].cell[%d].freq=%f",
		idx, ch.Index.Index, cell.Index.Index, *cmd.State,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell frequency: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"bus %d eq channel %d cell %d frequency set: %f\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State,
	)
	return nil
}

// BusEqCellGainCmd defines the command for getting or setting the gain of a specific EQ channel cell of a bus.
type BusEqCellGainCmd struct {
	State *float64
}

func (cmd *BusEqCellGainCmd) Run(
	ctx *appContext,
	bus *BusCmdGroup,
	ch *BusEqChannelCmd,
	cell *BusEqChannelCellCmd,
) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		return fmt.Errorf("bus eq channel cell gain queries are not supported")
	}

	txt := fmt.Sprintf(
		"bus[%d].eq.channel[%d].cell[%d].gain=%f",
		idx, ch.Index.Index, cell.Index.Index, *cmd.State,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell gain: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"bus %d eq channel %d cell %d gain set: %f\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State,
	)
	return nil
}

// BusEqCellQCmd defines the command for getting or setting the Q factor of a specific EQ channel cell of a bus.
type BusEqCellQCmd struct {
	State *float64
}

func (cmd *BusEqCellQCmd) Run(
	ctx *appContext,
	bus *BusCmdGroup,
	ch *BusEqChannelCmd,
	cell *BusEqChannelCellCmd,
) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		return fmt.Errorf("bus eq channel cell Q factor queries are not supported")
	}

	txt := fmt.Sprintf(
		"bus[%d].eq.channel[%d].cell[%d].q=%f",
		idx, ch.Index.Index, cell.Index.Index, *cmd.State,
	)

	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set EQ cell Q factor: %w", err)
	}

	fmt.Fprintf(
		ctx.Out,
		"bus %d eq channel %d cell %d Q factor set: %f\n",
		idx,
		ch.Index.Index,
		cell.Index.Index,
		*cmd.State,
	)
	return nil
}
