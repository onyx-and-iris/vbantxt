package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// BusCmdGroup defines the command group for controlling a specific bus, including its mute state and EQ settings.
type BusCmdGroup struct {
	Index struct {
		Index int         `arg:"" help:"The index of the bus. (0-based indexing)"`
		Mono  BusMonoCmd  `cmd:""       help:"Get or set the mono state of the bus."`
		Mute  BusMuteCmd  `cmd:""       help:"Get or set the mute state of the bus."`
		Mode  BusModeCmd  `cmd:""       help:"Get or set the mode of the bus."`
		Gain  BusGainCmd  `cmd:""       help:"Get or set the gain of the bus."`
		Label BusLabelCmd `cmd:""       help:"Get or set the label of the bus."`

		Eq BusEqCmdGroup `cmd:"" help:"Get or set the EQ settings of the bus."`
	} `arg:"" help:"Control a specific bus by index."`
}

// BusMonoCmd defines the command for getting or setting the mono state of a bus.
type BusMonoCmd struct {
	State *string `arg:"" help:"The mono state to set (true or false). If not provided, the current mono state will be returned." optional:"" enum:"true,false"`
}

// Run executes the BusMonoCmd command, either retrieving the current mono state of the bus or setting it based on the provided argument.
func (cmd *BusMonoCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Bus[idx].Mono()
		fmt.Fprintf(ctx.Out, "bus %d mono: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("bus[%d].mono=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set mono state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d mono set: %t\n", idx, *cmd.State == "true")
	return nil
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

	txt := fmt.Sprintf("bus[%d].mute=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d mute set: %t\n", idx, *cmd.State == "true")
	return nil
}

// BusModeCmd defines the command for getting or setting the mode of a bus.
type BusModeCmd struct {
	State *string `arg:"" help:"The mode to set for the bus. If not provided, the current mode will be returned." optional:"" enum:"normal,amix,bmix,repeat,composite,tvmix,upmix21,upmix41,upmix61,centeronly,lfeonly,rearonly"`
}

// Run executes the BusModeCmd command, either retrieving the current mode of the bus or setting it based on the provided argument.
func (cmd *BusModeCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		modeIndex := p.States().Bus[idx].Mode()

		candidates := []string{
			"normal", "amix", "bmix", "repeat", "composite", "tvmix",
			"upmix21", "upmix41", "upmix61", "centeronly", "lfeonly", "rearonly",
		}

		fmt.Fprintf(ctx.Out, "bus %d mode: %s\n", idx, candidates[modeIndex])
		return nil
	}

	txt := fmt.Sprintf("bus[%d].mode.%s=1", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set mode: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d mode set: %s\n", idx, *cmd.State)
	return nil
}

// BusGainCmd defines the command for getting or setting the gain of a bus.
type BusGainCmd struct {
	State *float64 `arg:"" help:"The gain to set for the bus. If not provided, the current gain will be returned." optional:""`
}

// Run executes the BusGainCmd command, either retrieving the current gain of the bus or setting it based on the provided argument.
func (cmd *BusGainCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		gain := p.BusGain[idx]
		fmt.Fprintf(ctx.Out, "bus %d gain: %f\n", idx, gain)
		return nil
	}

	txt := fmt.Sprintf("bus[%d].gain=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d gain set: %f\n", idx, *cmd.State)
	return nil
}

// BusLabelCmd defines the command for getting or setting the label of a bus.
type BusLabelCmd struct {
	State *string `arg:"" help:"The label to set for the bus. If not provided, the current label will be returned." optional:""`
}

// Run executes the BusLabelCmd command, either retrieving the current label of the bus or setting it based on the provided argument.
func (cmd *BusLabelCmd) Run(ctx *appContext, bus *BusCmdGroup) error {
	idx := bus.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		label := p.Labels().Bus[idx]
		fmt.Fprintf(ctx.Out, "bus %d label: %s\n", idx, label)
		return nil
	}

	txt := fmt.Sprintf("bus[%d].label=\"%s\"", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set label: %w", err)
	}

	fmt.Fprintf(ctx.Out, "bus %d label set: %s\n", idx, *cmd.State)
	return nil
}
