package packet

import (
	"bytes"
)

// VbanRTPacketNBS0 represents the body of an RTPacket with ident:0.
type VbanRTPacketNBS0 struct {
	Hdr VbanRTPacketHeader

	InputLeveldB100  []byte // 44–112
	OutputLeveldB100 []byte // 112–240
	TransportBit     []byte // 240–244
	StripState       []byte // 244–276
	BusState         []byte // 276–308

	StripGainLayer [8][8]float64
	StripGain      [8]float64

	BusGain [8]float64

	StripLabelUTF8c60 []byte // 452–932
	BusLabelUTF8c60   []byte // 932–1412
}

func (p *VbanRTPacketNBS0) NBS() NBS                    { return p.Hdr.NBSValue }
func (p *VbanRTPacketNBS0) Header() *VbanRTPacketHeader { return &p.Hdr }

func (p *VbanRTPacketNBS0) Equals(other *VbanRTPacketNBS0) bool {
	if other == nil {
		return false
	}

	// StripState
	if !bytes.Equal(p.StripState, other.StripState) {
		return false
	}

	// BusState
	if !bytes.Equal(p.BusState, other.BusState) {
		return false
	}

	// StripGainLayer (8 layers × 8 strips)
	for layer := 0; layer < 8; layer++ {
		for strip := 0; strip < 8; strip++ {
			if p.StripGainLayer[layer][strip] != other.StripGainLayer[layer][strip] {
				return false
			}
		}
	}

	// BusGain (8 floats)
	for i := 0; i < 8; i++ {
		if p.BusGain[i] != other.BusGain[i] {
			return false
		}
	}

	// StripLabelUTF8c60
	if !bytes.Equal(p.StripLabelUTF8c60, other.StripLabelUTF8c60) {
		return false
	}

	// BusLabelUTF8c60
	if !bytes.Equal(p.BusLabelUTF8c60, other.BusLabelUTF8c60) {
		return false
	}

	return true
}

func (p *VbanRTPacketNBS0) LevelsEqual(other *VbanRTPacketNBS0) bool {
	if other == nil {
		return false
	}

	// Strip levels (34)
	aStrip := p.StripLevels()
	bStrip := other.StripLevels()
	if len(aStrip) != len(bStrip) {
		return false
	}
	for i := range aStrip {
		if aStrip[i] != bStrip[i] {
			return false
		}
	}

	// Bus levels (64)
	aBus := p.BusLevels()
	bBus := other.BusLevels()
	if len(aBus) != len(bBus) {
		return false
	}
	for i := range aBus {
		if aBus[i] != bBus[i] {
			return false
		}
	}

	return true
}

func (p *VbanRTPacketNBS0) StripLevels() []float64 {
	raw := p.InputLeveldB100 // 44–112 bytes
	if len(raw) < 34*2 {
		return nil
	}

	levels := make([]float64, 34)
	for i := 0; i < 34; i++ {
		v := int16(raw[i*2]) | int16(raw[i*2+1])<<8
		levels[i] = float64(v) * 0.01
	}
	return levels
}

func (p *VbanRTPacketNBS0) BusLevels() []float64 {
	raw := p.OutputLeveldB100 // 112–240 bytes
	if len(raw) < 64*2 {
		return nil
	}

	levels := make([]float64, 64)
	for i := 0; i < 64; i++ {
		v := int16(raw[i*2]) | int16(raw[i*2+1])<<8
		levels[i] = float64(v) * 0.01
	}
	return levels
}

type Levels struct {
	Strip []float64
	Bus   []float64
}

func (p *VbanRTPacketNBS0) Levels() Levels {
	return Levels{
		Strip: p.StripLevels(),
		Bus:   p.BusLevels(),
	}
}

type States struct {
	Strip []ChannelState
	Bus   []ChannelState
}

func (p *VbanRTPacketNBS0) States() States {
	strip := make([]ChannelState, 8)
	bus := make([]ChannelState, 8)

	for i := 0; i < 8; i++ {
		strip[i] = NewChannelState(p.StripState[i*4 : i*4+4])
		bus[i] = NewChannelState(p.BusState[i*4 : i*4+4])
	}

	return States{Strip: strip, Bus: bus}
}

type Labels struct {
	Strip []string
	Bus   []string
}

func extractLabels(b []byte) []string {
	out := make([]string, len(b)/60)
	for i := 0; i < len(out); i++ {
		chunk := b[i*60 : (i+1)*60]
		n := bytes.IndexByte(chunk, 0)
		if n < 0 {
			n = len(chunk)
		}
		out[i] = string(chunk[:n])
	}
	return out
}

func (p *VbanRTPacketNBS0) Labels() Labels {
	return Labels{
		Strip: extractLabels(p.StripLabelUTF8c60),
		Bus:   extractLabels(p.BusLabelUTF8c60),
	}
}

func ParseVbanRTPacketNBS0(nbs NBS, data []byte) (*VbanRTPacketNBS0, error) {
	p := &VbanRTPacketNBS0{
		Hdr:               parseHeader(nbs, data),
		InputLeveldB100:   data[44:112],
		OutputLeveldB100:  data[112:240],
		TransportBit:      data[240:244],
		StripState:        append([]byte(nil), data[244:276]...),
		BusState:          append([]byte(nil), data[276:308]...),
		StripLabelUTF8c60: append([]byte(nil), data[452:932]...),
		BusLabelUTF8c60:   append([]byte(nil), data[932:1412]...),
	}

	const (
		layer1Offset = 308
		layerSize    = 16
	)

	for layer := range 8 {
		offset := layer1Offset + layer*layerSize
		raw := data[offset : offset+layerSize]

		for strip := range 8 {
			v := int16(raw[strip*2]) | int16(raw[strip*2+1])<<8
			p.StripGainLayer[layer][strip] = float64(v) * 0.01
		}
	}

	for strip := range 8 {
		maxVal := p.StripGainLayer[0][strip]
		for layer := 1; layer < 8; layer++ {
			if p.StripGainLayer[layer][strip] > maxVal {
				maxVal = p.StripGainLayer[layer][strip]
			}
		}
		p.StripGain[strip] = maxVal
	}

	rawBusGain := data[436:452]
	for bus := range 8 {
		v := int16(rawBusGain[bus*2]) | int16(rawBusGain[bus*2+1])<<8
		p.BusGain[bus] = float64(v) * 0.01
	}

	return p, nil
}
