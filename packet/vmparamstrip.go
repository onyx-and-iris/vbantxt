package packet

import (
	"encoding/binary"
	"math"
)

const VMPARAMSTRIP_SIZE = 174

type VbanVMParamStrip struct {
	Mode       uint32
	DBLevel    float64
	Audibility float64

	Pos3DX    float64
	Pos3DY    float64
	PosColorX float64
	PosColorY float64

	EQGain1 float64
	EQGain2 float64
	EQGain3 float64

	PEQ_On   [6]bool
	PEQ_Type [6]uint8
	PEQ_Gain [6]float64
	PEQ_Freq [6]float64
	PEQ_Q    [6]float64

	AudibilityC float64
	AudibilityG float64
	AudibilityD float64

	PosModX float64
	PosModY float64

	SendReverb float64
	SendDelay  float64
	SendFX1    float64
	SendFX2    float64

	DBLimit float64
	Karaoke float64

	CompGainIn    float64
	CompAttackMS  float64
	CompReleaseMS float64
	CompNKnee     float64
	CompRate      float64
	CompThreshold float64
	CompEnabled   bool
	CompAuto      bool
	CompGainOut   float64

	GateThreshold float64
	GateDamping   float64
	GateSidechain float64
	GateAttackMS  float64
	GateHoldMS    float64
	GateReleaseMS float64

	DenoiserThreshold float64

	PitchEnabled     bool
	PitchDryWet      float64
	PitchValue       float64
	PitchFormantLo   float64
	PitchFormantMed  float64
	PitchFormantHigh float64
}

func f16(b0, b1 byte) float64 {
	v := int16(b0) | int16(b1)<<8
	return float64(v) * 0.01
}

func f32(b []byte) float64 {
	return float64(math.Float32frombits(binary.LittleEndian.Uint32(b)))
}

func ParseVbanVMParamStrip(data []byte) VbanVMParamStrip {
	p := VbanVMParamStrip{}

	p.Mode = binary.LittleEndian.Uint32(data[0:4])
	p.DBLevel = f16(data[4], data[5]) // Python treats dblevel as float
	p.Audibility = f16(data[8], data[9])

	p.Pos3DX = f16(data[10], data[11])
	p.Pos3DY = f16(data[12], data[13])
	p.PosColorX = f16(data[14], data[15])
	p.PosColorY = f16(data[16], data[17])

	p.EQGain1 = f16(data[18], data[19])
	p.EQGain2 = f16(data[20], data[21])
	p.EQGain3 = f16(data[22], data[23])

	// PEQ On (6 bytes)
	for i := range 6 {
		p.PEQ_On[i] = data[24+i] != 0
	}

	// PEQ Type (6 bytes)
	copy(p.PEQ_Type[:], data[30:36])

	// PEQ Gain (6 floats)
	for i := range 6 {
		p.PEQ_Gain[i] = f32(data[36+i*4 : 40+i*4])
		p.PEQ_Freq[i] = f32(data[60+i*4 : 64+i*4])
		p.PEQ_Q[i] = f32(data[84+i*4 : 88+i*4])
	}

	p.AudibilityC = f16(data[108], data[109])
	p.AudibilityG = f16(data[110], data[111])
	p.AudibilityD = f16(data[112], data[113])

	p.PosModX = f16(data[114], data[115])
	p.PosModY = f16(data[116], data[117])

	p.SendReverb = f16(data[118], data[119])
	p.SendDelay = f16(data[120], data[121])
	p.SendFX1 = f16(data[122], data[123])
	p.SendFX2 = f16(data[124], data[125])

	p.DBLimit = f16(data[126], data[127])
	p.Karaoke = f16(data[128], data[129])

	p.CompGainIn = f16(data[130], data[131])
	p.CompAttackMS = f16(data[132], data[133])
	p.CompReleaseMS = f16(data[134], data[135])
	p.CompNKnee = f16(data[136], data[137])
	p.CompRate = f16(data[138], data[139])
	p.CompThreshold = f16(data[140], data[141])

	p.CompEnabled = data[142] != 0
	p.CompAuto = data[144] != 0

	p.CompGainOut = f16(data[146], data[147])

	p.GateThreshold = f16(data[148], data[149])
	p.GateDamping = f16(data[150], data[151])
	p.GateSidechain = f16(data[152], data[153])
	p.GateAttackMS = f16(data[154], data[155])
	p.GateHoldMS = f16(data[156], data[157])
	p.GateReleaseMS = f16(data[158], data[159])

	p.DenoiserThreshold = f16(data[160], data[161])

	p.PitchEnabled = data[162] != 0
	p.PitchDryWet = f16(data[164], data[165])
	p.PitchValue = f16(data[166], data[167])
	p.PitchFormantLo = f16(data[168], data[169])
	p.PitchFormantMed = f16(data[170], data[171])
	p.PitchFormantHigh = f16(data[172], data[173])

	return p
}
