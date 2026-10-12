package packet

import (
	"encoding/binary"
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/logging"
)

type RTPacket interface {
	NBS() NBS
	Header() *VbanRTPacketHeader
}

type VbanRTPacketHeader struct {
	NBSValue           NBS
	VoicemeeterType    uint8
	Reserved           uint8
	BufferSize         uint16
	VoicemeeterVersion [4]byte
	OptionBits         [4]byte
	SampleRate         uint32
}

func parseHeader(nbs NBS, data []byte) VbanRTPacketHeader {
	h := VbanRTPacketHeader{
		NBSValue:        nbs,
		VoicemeeterType: data[28],
		Reserved:        data[29],
		BufferSize:      binary.LittleEndian.Uint16(data[30:32]),
		SampleRate:      binary.LittleEndian.Uint32(data[40:44]),
	}

	copy(h.VoicemeeterVersion[:], data[32:36])
	copy(h.OptionBits[:], data[36:40])

	return h
}

func ParseRTPacket(data []byte) (RTPacket, error) {
	nbs := NBS(data[5])
	logging.Tracef("Parsing RT packet with NBS type: %d", nbs)

	switch nbs {
	case NBSZero:
		return ParseVbanRTPacketNBS0(nbs, data)
	case NBSOne:
		return ParseVbanRTPacketNBS1(nbs, data)
	default:
		return nil, fmt.Errorf("unsupported NBS type: %d", nbs)
	}
}
