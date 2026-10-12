package packet

// ChannelModes is a bitmask enum matching Voicemeeter's internal flags.
const (
	// Common boolean modes
	ChannelModeMute uint32 = 0x00000001
	ChannelModeSolo uint32 = 0x00000002
	ChannelModeMono uint32 = 0x00000004
	ChannelModeMC   uint32 = 0x00000008

	// Mix modes
	ChannelModeAMix       uint32 = 0x00000010
	ChannelModeRepeat     uint32 = 0x00000020
	ChannelModeBMix       uint32 = 0x00000030
	ChannelModeComposite  uint32 = 0x00000040
	ChannelModeTVMix      uint32 = 0x00000050
	ChannelModeUpmix21    uint32 = 0x00000060
	ChannelModeUpmix41    uint32 = 0x00000070
	ChannelModeUpmix61    uint32 = 0x00000080
	ChannelModeCenterOnly uint32 = 0x00000090
	ChannelModeLFEOnly    uint32 = 0x000000A0
	ChannelModeRearOnly   uint32 = 0x000000B0

	ChannelModeMask uint32 = 0x000000F0

	// EQ modes
	ChannelModeEQOn  uint32 = 0x00000100
	ChannelModeCross uint32 = 0x00000200
	ChannelModeEQAB  uint32 = 0x00000800

	// Bus assignments (strip → bus routing)
	ChannelModeBusA  uint32 = 0x00001000
	ChannelModeBusA1 uint32 = 0x00001000
	ChannelModeBusA2 uint32 = 0x00002000
	ChannelModeBusA3 uint32 = 0x00004000
	ChannelModeBusA4 uint32 = 0x00008000
	ChannelModeBusA5 uint32 = 0x00080000

	ChannelModeBusB  uint32 = 0x00010000
	ChannelModeBusB1 uint32 = 0x00010000
	ChannelModeBusB2 uint32 = 0x00020000
	ChannelModeBusB3 uint32 = 0x00040000

	// Pan modes
	ChannelModePan0     uint32 = 0x00000000
	ChannelModePanColor uint32 = 0x00100000
	ChannelModePanMod   uint32 = 0x00200000
	ChannelModePanMask  uint32 = 0x00F00000

	// Post‑FX
	ChannelModePostFXR uint32 = 0x01000000
	ChannelModePostFXD uint32 = 0x02000000
	ChannelModePostFX1 uint32 = 0x04000000
	ChannelModePostFX2 uint32 = 0x08000000

	// Selection / monitoring
	ChannelModeSel     uint32 = 0x10000000
	ChannelModeMonitor uint32 = 0x20000000
)
