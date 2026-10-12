package main

import (
	"context"
	"fmt"

	"github.com/onyx-and-iris/vbantxt/packet"
)

type GainLayerCmdGroup struct {
	Index struct {
		Index int               `arg:"" help:"Index of the gain layer."`
		Level GainLayerLevelCmd `cmd:"" help:"Control the level of the gain layer."`
	} `arg:"" help:"Control a specific gain layer by its index."`
}

type GainLayerLevelCmd struct {
	Level *float64 `arg:"" optional:"" help:""`
}

func (cmd *GainLayerLevelCmd) Run(
	ctx *appContext,
	strip *StripCmdGroup,
	gainlayer *GainLayerCmdGroup,
) error {
	idx_gainlayer := gainlayer.Index.Index
	idx_strip := strip.Index.Index

	if cmd.Level == nil {
		opCtx, cancel := context.WithTimeout(ctx.Ctx, ctx.Timeout)
		defer cancel()

		pkt, err := ctx.Client.WaitForRT(opCtx, packet.NBSZero)
		if err != nil {
			return err
		}

		p := pkt.(*packet.VbanRTPacketNBS0)
		v := p.StripGainLayer[idx_strip][idx_gainlayer]
		fmt.Fprintf(
			ctx.Out,
			"strip %d gain layer %d level: %f\n",
			idx_strip,
			idx_gainlayer,
			v,
		)
		return nil
	}

	txt := fmt.Sprintf("strip[%d].gainlayer[%d]=%f", idx_strip, idx_gainlayer, *cmd.Level)
	if err := ctx.Client.SendTXT(ctx.Ctx, txt); err != nil {
		return fmt.Errorf("failed to set gain layer level: %w", err)
	}

	return nil
}
