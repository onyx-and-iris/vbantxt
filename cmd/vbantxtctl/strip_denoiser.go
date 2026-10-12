package main

import (
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

type StripDenoiserCmdGroup struct {
	Knob      StripDenoiserKnobCmd      `cmd:"" help:"Get or set the knob settings of the denoiser."`
	Threshold StripDenoiserThresholdCmd `cmd:"" help:"Get or set the threshold of the denoiser."`
}

type StripDenoiserKnobCmd struct {
	State *float64 `arg:"" optional:"" help:"Knob setting of the denoiser."`
}

func (cmd *StripDenoiserKnobCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].AudibilityD

		fmt.Fprintf(ctx.Out, "strip %d denoiser knob: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].denoiser=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set denoiser knob: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d denoiser knob set: %f\n", idx, *cmd.State)
	return nil
}

type StripDenoiserThresholdCmd struct {
	State *float64 `arg:"" optional:"" help:"Threshold of the denoiser."`
}

func (cmd *StripDenoiserThresholdCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		pkt, err := ctx.Client.WaitForRT(ctx.Ctx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].DenoiserThreshold

		fmt.Fprintf(ctx.Out, "strip %d denoiser threshold: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].denoiser.threshold=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set denoiser threshold: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d denoiser threshold set: %f\n", idx, *cmd.State)
	return nil
}
