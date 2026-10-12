package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/packet"
)

type StripCompCmdGroup struct {
	Knob       StripCompKnobCmd       `cmd:"" help:"Get or set the knob settings of the compressor."`
	InputGain  StripCompInputGainCmd  `cmd:"" help:"Get or set the input gain of the compressor."`
	Ratio      StripCompRatioCmd      `cmd:"" help:"Get or set the ratio of the compressor."`
	Threshold  StripCompThresholdCmd  `cmd:"" help:"Get or set the threshold of the compressor."`
	Attack     StripCompAttackCmd     `cmd:"" help:"Get or set the attack of the compressor."`
	Release    StripCompReleaseCmd    `cmd:"" help:"Get or set the release of the compressor."`
	Knee       StripCompKneeCmd       `cmd:"" help:"Get or set the knee of the compressor."`
	OutputGain StripCompOutputGainCmd `cmd:"" help:"Get or set the output gain of the compressor."`
	Makeup     StripCompMakeupCmd     `cmd:"" help:"Get or set the makeup gain of the compressor."`
}

type StripCompKnobCmd struct {
	State *float64 `arg:"" optional:"" help:"Knob setting of the compressor."`
}

func (cmd *StripCompKnobCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].AudibilityC

		fmt.Fprintf(ctx.Out, "strip %d comp knob: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp knob: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp knob set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompInputGainCmd struct {
	State *float64 `arg:"" optional:"" help:"Input gain of the compressor."`
}

func (cmd *StripCompInputGainCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompGainIn

		fmt.Fprintf(ctx.Out, "strip %d comp input gain: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.gainin=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp input gain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp input gain set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompRatioCmd struct {
	State *float64 `arg:"" optional:"" help:"Ratio of the compressor."`
}

func (cmd *StripCompRatioCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompRate

		fmt.Fprintf(ctx.Out, "strip %d comp ratio: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.ratio=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp ratio: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp ratio set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompThresholdCmd struct {
	State *float64 `arg:"" optional:"" help:"Threshold of the compressor."`
}

func (cmd *StripCompThresholdCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompThreshold

		fmt.Fprintf(ctx.Out, "strip %d comp threshold: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.threshold=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp threshold: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp threshold set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompAttackCmd struct {
	State *float64 `arg:"" optional:"" help:"Attack of the compressor."`
}

func (cmd *StripCompAttackCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompAttackMS

		fmt.Fprintf(ctx.Out, "strip %d comp attack: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.attack=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp attack: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp attack set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompReleaseCmd struct {
	State *float64 `arg:"" optional:"" help:"Release of the compressor."`
}

func (cmd *StripCompReleaseCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompReleaseMS

		fmt.Fprintf(ctx.Out, "strip %d comp release: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.release=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp release: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp release set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompKneeCmd struct {
	State *float64 `arg:"" optional:"" help:"Knee of the compressor."`
}

func (cmd *StripCompKneeCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompNKnee

		fmt.Fprintf(ctx.Out, "strip %d comp knee: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.knee=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp knee: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp knee set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompOutputGainCmd struct {
	State *float64 `arg:"" optional:"" help:"Output gain of the compressor."`
}

func (cmd *StripCompOutputGainCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompGainOut

		fmt.Fprintf(ctx.Out, "strip %d comp output gain: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.gainout=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp output gain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp output gain set: %f\n", idx, *cmd.State)
	return nil
}

type StripCompMakeupCmd struct {
	State *bool `arg:"" optional:"" help:"Enable or disable automatic makeup gain of the compressor."`
}

func (cmd *StripCompMakeupCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].CompAuto

		fmt.Fprintf(ctx.Out, "strip %d comp makeup gain: %t\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].comp.makeup=%t", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set comp makeup gain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d comp makeup gain set: %t\n", idx, *cmd.State)
	return nil
}
