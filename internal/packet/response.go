package packet

import (
	"encoding/binary"
	"fmt"

	"github.com/onyx-and-iris/vbantxt/internal/logging"
)

type VbanRTResponseHeader struct {
	Name         string
	FormatSR     uint8
	FormatNBS    uint8
	FormatNBC    uint8
	FormatBIT    uint8
	FrameCounter uint32
}

func ParseVbanRTResponseHeader(data []byte) (*VbanRTResponseHeader, error) {
	if len(data) < HeaderSize {
		return nil, fmt.Errorf("response too short: %d bytes", len(data))
	}

	if string(data[0:4]) != "VBAN" {
		return nil, fmt.Errorf("not a VBAN packet")
	}

	h := &VbanRTResponseHeader{
		FormatSR:  data[4],
		FormatNBS: data[5],
		FormatNBC: data[6],
		FormatBIT: data[7],
	}
	logging.Tracef("Raw header bytes: %v", data[:8])
	logging.Tracef("SR=%d NBS=%d NBC=%d BIT=%d", h.FormatSR, h.FormatNBS, h.FormatNBC, h.FormatBIT)

	// Validate this is an RTPACKET response
	if h.FormatNBC != uint8(ServiceRTPacket) {
		return nil, fmt.Errorf("not an RTPACKET response packet: 0x%02x", h.FormatNBC)
	}

	// Extract stream name (null‑terminated)
	rawName := data[8:24]
	n := 0
	for ; n < len(rawName); n++ {
		if rawName[n] == 0 {
			break
		}
	}
	h.Name = string(rawName[:n])

	// Framecounter (LE)
	h.FrameCounter = binary.LittleEndian.Uint32(data[24:28])

	return h, nil
}
