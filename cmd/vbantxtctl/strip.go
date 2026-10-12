package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// StripCmdGroup defines the command group for controlling a specific strip, including its mute state and EQ settings.
type StripCmdGroup struct {
	Index struct {
		Index int           `arg:"" help:"The index of the strip. (0-based indexing)"`
		Mono  StripMonoCmd  `cmd:""       help:"Get or set the mono state of the strip."`
		Solo  StripSoloCmd  `cmd:""       help:"Get or set the solo state of the strip."`
		Mute  StripMuteCmd  `cmd:""       help:"Get or set the mute state of the strip."`
		Gain  StripGainCmd  `cmd:""       help:"Get or set the gain of the strip."`
		Label StripLabelCmd `cmd:""       help:"Get or set the label of the strip."`

		A1 StripA1Cmd `cmd:""       help:"Get or set the A1 state of the strip." name:"a1"`
		A2 StripA2Cmd `cmd:""       help:"Get or set the A2 state of the strip." name:"a2"`
		A3 StripA3Cmd `cmd:""       help:"Get or set the A3 state of the strip." name:"a3"`
		A4 StripA4Cmd `cmd:""       help:"Get or set the A4 state of the strip." name:"a4"`
		A5 StripA5Cmd `cmd:""       help:"Get or set the A5 state of the strip." name:"a5"`
		B1 StripB1Cmd `cmd:""       help:"Get or set the B1 state of the strip." name:"b1"`
		B2 StripB2Cmd `cmd:""       help:"Get or set the B2 state of the strip." name:"b2"`
		B3 StripB3Cmd `cmd:""       help:"Get or set the B3 state of the strip." name:"b3"`

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

	txt := fmt.Sprintf("strip[%d].mono=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
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

	txt := fmt.Sprintf("strip[%d].solo=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
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

	txt := fmt.Sprintf("strip[%d].mute=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set mute state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d mute set: %t\n", idx, *cmd.State == "true")
	return nil
}

type StripGainCmd struct {
	State *float64 `arg:"" help:"The gain to set for the strip. If not provided, the current gain will be returned." optional:""`
}

// Run executes the StripGainCmd command, either retrieving the current gain of the strip or setting it based on the provided argument.
func (cmd *StripGainCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		gain := p.StripGain[idx]
		fmt.Fprintf(ctx.Out, "strip %d gain: %f\n", idx, gain)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gain=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gain set: %f\n", idx, *cmd.State)
	return nil
}

// StripA1Cmd defines the command for getting or setting the A1 state of a strip.
type StripA1Cmd struct {
	State *string `arg:"" help:"The A1 state to set (true or false). If not provided, the current A1 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripA1Cmd command, either retrieving the current A1 state of the strip or setting it based on the provided argument.
func (cmd *StripA1Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusA1()
		fmt.Fprintf(ctx.Out, "strip %d A1: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].A1=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set A1 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d A1 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripA2Cmd defines the command for getting or setting the A2 state of a strip.
type StripA2Cmd struct {
	State *string `arg:"" help:"The A2 state to set (true or false). If not provided, the current A2 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripA2Cmd command, either retrieving the current A2 state of the strip or setting it based on the provided argument.
func (cmd *StripA2Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusA2()
		fmt.Fprintf(ctx.Out, "strip %d A2: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].A2=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set A2 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d A2 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripA3Cmd defines the command for getting or setting the A3 state of a strip.
type StripA3Cmd struct {
	State *string `arg:"" help:"The A3 state to set (true or false). If not provided, the current A3 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripA3Cmd command, either retrieving the current A3 state of the strip or setting it based on the provided argument.
func (cmd *StripA3Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusA3()
		fmt.Fprintf(ctx.Out, "strip %d A3: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].A3=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set A3 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d A3 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripA4Cmd defines the command for getting or setting the A4 state of a strip.
type StripA4Cmd struct {
	State *string `arg:"" help:"The A4 state to set (true or false). If not provided, the current A4 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripA4Cmd command, either retrieving the current A4 state of the strip or setting it based on the provided argument.
func (cmd *StripA4Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusA4()
		fmt.Fprintf(ctx.Out, "strip %d A4: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].A4=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set A4 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d A4 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripA5Cmd defines the command for getting or setting the A5 state of a strip.
type StripA5Cmd struct {
	State *string `arg:"" help:"The A5 state to set (true or false). If not provided, the current A5 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripA5Cmd command, either retrieving the current A5 state of the strip or setting it based on the provided argument.
func (cmd *StripA5Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusA5()
		fmt.Fprintf(ctx.Out, "strip %d A5: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].A5=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set A5 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d A5 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripB1Cmd defines the command for getting or setting the B1 state of a strip.
type StripB1Cmd struct {
	State *string `arg:"" help:"The B1 state to set (true or false). If not provided, the current B1 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripB1Cmd command, either retrieving the current B1 state of the strip or setting it based on the provided argument.
func (cmd *StripB1Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusB1()
		fmt.Fprintf(ctx.Out, "strip %d B1: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].B1=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set B1 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d B1 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripB2Cmd defines the command for getting or setting the B2 state of a strip.
type StripB2Cmd struct {
	State *string `arg:"" help:"The B2 state to set (true or false). If not provided, the current B2 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripB2Cmd command, either retrieving the current B2 state of the strip or setting it based on the provided argument.
func (cmd *StripB2Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusB2()
		fmt.Fprintf(ctx.Out, "strip %d B2: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].B2=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set B2 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d B2 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripB3Cmd defines the command for getting or setting the B3 state of a strip.
type StripB3Cmd struct {
	State *string `arg:"" help:"The B3 state to set (true or false). If not provided, the current B3 state will be returned." optional:"" enum:"true,false"`
}

// Run executes the StripB3Cmd command, either retrieving the current B3 state of the strip or setting it based on the provided argument.
func (cmd *StripB3Cmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.States().Strip[idx].BusB3()
		fmt.Fprintf(ctx.Out, "strip %d B3: %t\n", idx, v)
		return nil
	}

	var val int
	if *cmd.State == "true" {
		val = 1
	}

	txt := fmt.Sprintf("strip[%d].B3=%d", idx, val)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set B3 state: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d B3 set: %t\n", idx, *cmd.State == "true")
	return nil
}

// StripLabelCmd defines the command for getting or setting the label of a strip.
type StripLabelCmd struct {
	Label *string `arg:"" help:"The label to set for the strip. If not provided, the current label will be returned." optional:""`
}

// Run executes the StripLabelCmd command, either retrieving the current label of the strip or setting it based on the provided argument.
func (cmd *StripLabelCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.Label == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.Labels().Strip[idx]
		fmt.Fprintf(ctx.Out, "strip %d label: %s\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].Label=\"%s\"", idx, *cmd.Label)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set label: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d label set: %s\n", idx, *cmd.Label)
	return nil
}
