package main

import (
	"context"
	"fmt"

	"github.com/onyx-and-iris/vbantxt/packet"
)

type StripGateCmdGroup struct {
	Knob        StripGateKnobCmd        `cmd:"" help:"Get or set the knob settings of the gate."`
	Threshold   StripGateThresholdCmd   `cmd:"" help:"Get or set the threshold of the gate."`
	DampingMax  StripGateDampingMaxCmd  `cmd:"" help:"Get or set the maximum damping of the gate."`
	BPSidechain StripGateBPSidechainCmd `cmd:"" help:"Get or set the sidechain of the gate."`
	Attack      StripGateAttackCmd      `cmd:"" help:"Get or set the attack of the gate."`
	Hold        StripGateHoldCmd        `cmd:"" help:"Get or set the hold of the gate."`
	Release     StripGateReleaseCmd     `cmd:"" help:"Get or set the release of the gate."`
}

type StripGateKnobCmd struct {
	State *float64 `arg:"" optional:"" help:"Knob setting of the gate."`
}

func (cmd *StripGateKnobCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].AudibilityG

		fmt.Fprintf(ctx.Out, "strip %d gate knob: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate knob: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate knob set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateThresholdCmd struct {
	State *float64 `arg:"" optional:"" help:"Threshold of the gate."`
}

func (cmd *StripGateThresholdCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateThreshold

		fmt.Fprintf(ctx.Out, "strip %d gate threshold: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.threshold=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate threshold: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate threshold set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateDampingMaxCmd struct {
	State *float64 `arg:"" optional:"" help:"Maximum damping of the gate."`
}

func (cmd *StripGateDampingMaxCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateDamping

		fmt.Fprintf(ctx.Out, "strip %d gate maximum damping: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.damping=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate maximum damping: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate maximum damping set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateBPSidechainCmd struct {
	State *float64 `arg:"" optional:"" help:"Sidechain of the gate."`
}

func (cmd *StripGateBPSidechainCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateSidechain

		fmt.Fprintf(ctx.Out, "strip %d gate sidechain: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.bpsidechain=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate sidechain: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate sidechain set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateAttackCmd struct {
	State *float64 `arg:"" optional:"" help:"Attack of the gate."`
}

func (cmd *StripGateAttackCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateAttackMS

		fmt.Fprintf(ctx.Out, "strip %d gate attack: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.attack=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate attack: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate attack set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateHoldCmd struct {
	State *float64 `arg:"" optional:"" help:"Hold of the gate."`
}

func (cmd *StripGateHoldCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateHoldMS

		fmt.Fprintf(ctx.Out, "strip %d gate hold: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.hold=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate hold: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate hold set: %f\n", idx, *cmd.State)
	return nil
}

type StripGateReleaseCmd struct {
	State *float64 `arg:"" optional:"" help:"Release of the gate."`
}

func (cmd *StripGateReleaseCmd) Run(ctx *appContext, strip *StripCmdGroup) error {
	idx := strip.Index.Index

	if cmd.State == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSOne)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS1)
		v := p.Strip[idx].GateReleaseMS

		fmt.Fprintf(ctx.Out, "strip %d gate release: %f\n", idx, v)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gate.release=%f", idx, *cmd.State)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gate release: %w", err)
	}

	fmt.Fprintf(ctx.Out, "strip %d gate release set: %f\n", idx, *cmd.State)
	return nil
}
