package main

type StripGateCmdGroup struct {
	Knob        StripGateKnobCmd        `cmd:"" help:"Get or set the knob settings of the gate."`
	Threshold   StripGateThresholdCmd   `cmd:"" help:"Get or set the threshold of the gate."`
	DampingMax  StripGateDampingMaxCmd  `cmd:"" help:"Get or set the maximum damping of the gate."`
	BPSidechain StripGateBPSidechainCmd `cmd:"" help:"Get or set the sidechain of the gate."`
	Attack      StripGateAttackCmd      `cmd:"" help:"Get or set the attack of the gate."`
	Hold        StripGateHoldCmd        `cmd:"" help:"Get or set the hold of the gate."`
	Release     StripGateReleaseCmd     `cmd:"" help:"Get or set the release of the gate."`
}

type StripGateKnobCmd struct{ State *float64 }

type StripGateThresholdCmd struct{ State *float64 }

type StripGateDampingMaxCmd struct{ State *float64 }

type StripGateBPSidechainCmd struct{ State *float64 }

type StripGateAttackCmd struct{ State *float64 }

type StripGateHoldCmd struct{ State *float64 }

type StripGateReleaseCmd struct{ State *float64 }
