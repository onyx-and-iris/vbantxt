package packet

import (
	"encoding/binary"
	"math"
)

const VMPARAMSTRIP_SIZE = 174

type VbanVMParamStrip struct {
	Mode       uint32
	DBLevel    int32
	Audibility int16
	Pos3DX     int16
	Pos3DY     int16
	PosColorX  int16
	PosColorY  int16
	EQGain1    int16
	EQGain2    int16
	EQGain3    int16

	PEQ_On   [6]uint8
	PEQ_Type [6]uint8
	PEQ_Gain [6]float32
	PEQ_Freq [6]float32
	PEQ_Q    [6]float32

	AudibilityC int16
	AudibilityG int16
	AudibilityD int16
	PosModX     int16
	PosModY     int16
	SendReverb  int16
	SendDelay   int16
	SendFX1     int16
	SendFX2     int16
	DBLimit     int16
	Karaoke     int16

	CompGainIn    int16
	CompAttackMS  int16
	CompReleaseMS int16
	CompNKnee     int16
	CompRate      int16
	CompThreshold int16
	CompEnabled   uint16
	CompAuto      uint16
	CompGainOut   int16

	GateThreshold int16
	GateDamping   int16
	GateSidechain uint16
	GateAttackMS  int16
	GateHoldMS    int16
	GateReleaseMS int16

	DenoiserThreshold int16
	PitchEnabled      uint16
	PitchDryWet       int16
	PitchValue        int16
	PitchFormantLo    int16
	PitchFormantMed   int16
	PitchFormantHigh  int16
}

func ParseVbanVMParamStrip(data []byte) VbanVMParamStrip {
	p := VbanVMParamStrip{}

	p.Mode = binary.LittleEndian.Uint32(data[0:4])
	p.DBLevel = int32(binary.LittleEndian.Uint32(data[4:8]))
	p.Audibility = int16(binary.LittleEndian.Uint16(data[8:10]))
	p.Pos3DX = int16(binary.LittleEndian.Uint16(data[10:12]))
	p.Pos3DY = int16(binary.LittleEndian.Uint16(data[12:14]))
	p.PosColorX = int16(binary.LittleEndian.Uint16(data[14:16]))
	p.PosColorY = int16(binary.LittleEndian.Uint16(data[16:18]))
	p.EQGain1 = int16(binary.LittleEndian.Uint16(data[18:20]))
	p.EQGain2 = int16(binary.LittleEndian.Uint16(data[20:22]))
	p.EQGain3 = int16(binary.LittleEndian.Uint16(data[22:24]))

	copy(p.PEQ_On[:], data[24:30])
	copy(p.PEQ_Type[:], data[30:36])

	for i := 0; i < 6; i++ {
		p.PEQ_Gain[i] = math.Float32frombits(binary.LittleEndian.Uint32(data[36+i*4 : 40+i*4]))
		p.PEQ_Freq[i] = math.Float32frombits(binary.LittleEndian.Uint32(data[60+i*4 : 64+i*4]))
		p.PEQ_Q[i] = math.Float32frombits(binary.LittleEndian.Uint32(data[84+i*4 : 88+i*4]))
	}

	p.AudibilityC = int16(binary.LittleEndian.Uint16(data[108:110]))
	p.AudibilityG = int16(binary.LittleEndian.Uint16(data[110:112]))
	p.AudibilityD = int16(binary.LittleEndian.Uint16(data[112:114]))
	p.PosModX = int16(binary.LittleEndian.Uint16(data[114:116]))
	p.PosModY = int16(binary.LittleEndian.Uint16(data[116:118]))
	p.SendReverb = int16(binary.LittleEndian.Uint16(data[118:120]))
	p.SendDelay = int16(binary.LittleEndian.Uint16(data[120:122]))
	p.SendFX1 = int16(binary.LittleEndian.Uint16(data[122:124]))
	p.SendFX2 = int16(binary.LittleEndian.Uint16(data[124:126]))
	p.DBLimit = int16(binary.LittleEndian.Uint16(data[126:128]))
	p.Karaoke = int16(binary.LittleEndian.Uint16(data[128:130]))

	p.CompGainIn = int16(binary.LittleEndian.Uint16(data[130:132]))
	p.CompAttackMS = int16(binary.LittleEndian.Uint16(data[132:134]))
	p.CompReleaseMS = int16(binary.LittleEndian.Uint16(data[134:136]))
	p.CompNKnee = int16(binary.LittleEndian.Uint16(data[136:138]))
	p.CompRate = int16(binary.LittleEndian.Uint16(data[138:140]))
	p.CompThreshold = int16(binary.LittleEndian.Uint16(data[140:142]))
	p.CompEnabled = binary.LittleEndian.Uint16(data[142:144])
	p.CompAuto = binary.LittleEndian.Uint16(data[144:146])
	p.CompGainOut = int16(binary.LittleEndian.Uint16(data[146:148]))

	p.GateThreshold = int16(binary.LittleEndian.Uint16(data[148:150]))
	p.GateDamping = int16(binary.LittleEndian.Uint16(data[150:152]))
	p.GateSidechain = binary.LittleEndian.Uint16(data[152:154])
	p.GateAttackMS = int16(binary.LittleEndian.Uint16(data[154:156]))
	p.GateHoldMS = int16(binary.LittleEndian.Uint16(data[156:158]))
	p.GateReleaseMS = int16(binary.LittleEndian.Uint16(data[158:160]))

	p.DenoiserThreshold = int16(binary.LittleEndian.Uint16(data[160:162]))
	p.PitchEnabled = binary.LittleEndian.Uint16(data[162:164])
	p.PitchDryWet = int16(binary.LittleEndian.Uint16(data[164:166]))
	p.PitchValue = int16(binary.LittleEndian.Uint16(data[166:168]))
	p.PitchFormantLo = int16(binary.LittleEndian.Uint16(data[168:170]))
	p.PitchFormantMed = int16(binary.LittleEndian.Uint16(data[170:172]))
	p.PitchFormantHigh = int16(binary.LittleEndian.Uint16(data[172:174]))

	return p
}
