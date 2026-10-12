package main

type StripDenoiserCmdGroup struct {
	Knob StripDenoiserKnobCmd `cmd:"" help:"Get or set the knob settings of the denoiser."`
}

type StripDenoiserKnobCmd struct{ State *float64 }
