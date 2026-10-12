package packet

import "encoding/binary"

type ChannelState struct {
	State uint32
}

func NewChannelState(b []byte) ChannelState {
	return ChannelState{
		State: binary.LittleEndian.Uint32(b),
	}
}

func (c ChannelState) Has(mode uint32) bool {
	return (c.State & mode) != 0
}

func (c ChannelState) Mute() bool {
	return (c.State & ChannelModeMute) != 0
}

func (c ChannelState) Solo() bool {
	return (c.State & ChannelModeSolo) != 0
}

func (c ChannelState) Mono() bool {
	return (c.State & ChannelModeMono) != 0
}

func (c ChannelState) MC() bool {
	return (c.State & ChannelModeMC) != 0
}

func (c ChannelState) EQOn() bool {
	return (c.State & ChannelModeEQOn) != 0
}

func (c ChannelState) EQAB() bool {
	return (c.State & ChannelModeEQAB) != 0
}

func (c ChannelState) BusA1() bool { return (c.State & ChannelModeBusA1) != 0 }
func (c ChannelState) BusA2() bool { return (c.State & ChannelModeBusA2) != 0 }
func (c ChannelState) BusA3() bool { return (c.State & ChannelModeBusA3) != 0 }
func (c ChannelState) BusA4() bool { return (c.State & ChannelModeBusA4) != 0 }
func (c ChannelState) BusA5() bool { return (c.State & ChannelModeBusA5) != 0 }

func (c ChannelState) BusB1() bool { return (c.State & ChannelModeBusB1) != 0 }
func (c ChannelState) BusB2() bool { return (c.State & ChannelModeBusB2) != 0 }
func (c ChannelState) BusB3() bool { return (c.State & ChannelModeBusB3) != 0 }

func (c ChannelState) PanColor() bool { return (c.State & ChannelModePanColor) != 0 }
func (c ChannelState) PanMod() bool   { return (c.State & ChannelModePanMod) != 0 }
func (c ChannelState) Sel() bool      { return (c.State & ChannelModeSel) != 0 }
func (c ChannelState) Monitor() bool  { return (c.State & ChannelModeMonitor) != 0 }
