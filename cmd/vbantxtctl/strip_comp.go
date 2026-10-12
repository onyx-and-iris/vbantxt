package main

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

type StripCompKnobCmd struct{ State *float64 }

type StripCompInputGainCmd struct{ State *float64 }

type StripCompRatioCmd struct{ State *float64 }

type StripCompThresholdCmd struct{ State *float64 }

type StripCompAttackCmd struct{ State *float64 }

type StripCompReleaseCmd struct{ State *float64 }

type StripCompKneeCmd struct{ State *float64 }

type StripCompOutputGainCmd struct{ State *float64 }

type StripCompMakeupCmd struct{ State *float64 }
