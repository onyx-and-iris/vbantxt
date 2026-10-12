package packet

import (
	"encoding/binary"
	"fmt"
)

// BpsOpts defines the available baud rate options.
var BpsOpts = []int{
	0, 110, 150, 300, 600, 1200, 2400, 4800, 9600, 14400, 19200, 31250,
	38400, 57600, 115200, 128000, 230400, 250000, 256000, 460800, 921600,
	1000000, 1500000, 2000000, 3000000,
}

// VbanRTRequestHeader represents the header of a VBAN-TXT request packet.
type VbanRTRequestHeader struct {
	Streamname string
	Channel    uint8
	BPSIndex   uint8
}

// NewVbanRTRequestHeader validates and constructs a VBAN-TXT header.
func NewVbanRTRequestHeader(streamname string) (*VbanRTRequestHeader, error) {
	if len(streamname) > StreamnameMaxLength {
		return nil, fmt.Errorf(
			"streamname too long: %d bytes, max %d",
			len(streamname),
			StreamnameMaxLength,
		)
	}

	bpsIndex, err := BpsIndex(256000)
	if err != nil {
		return nil, fmt.Errorf("default baud rate 256000 not found in BpsOpts")
	}

	return &VbanRTRequestHeader{
		Streamname: streamname,
		Channel:    0,
		BPSIndex:   bpsIndex,
	}, nil
}

func BpsIndex(bps int) (uint8, error) {
	for i, v := range BpsOpts {
		if v == bps {
			return uint8(i), nil
		}
	}
	return 0, fmt.Errorf("invalid BPS value %d", bps)
}

// streamname returns the 16-byte padded streamname field.
func (h *VbanRTRequestHeader) streamname() [StreamnameMaxLength]byte {
	var out [StreamnameMaxLength]byte
	copy(out[:], []byte(h.Streamname))
	return out
}

// ToBytes encodes the header into a byte slice suitable for transmission.
func (h *VbanRTRequestHeader) ToBytes(framecounter uint32) []byte {
	buf := make([]byte, 4+4+16+4) // 4s4B16sI

	copy(buf[0:4], []byte("VBAN"))

	buf[4] = h.BPSIndex | uint8(SubProtocolText)
	buf[5] = 0
	buf[6] = h.Channel
	buf[7] = uint8(StreamUTF8)

	sn := h.streamname()
	copy(buf[8:24], sn[:])

	binary.LittleEndian.PutUint32(buf[24:28], framecounter)

	return buf
}

// EncodeWithPayload builds header + UTF-8 payload.
func (h *VbanRTRequestHeader) EncodeWithPayload(framecounter uint32, payload string) []byte {
	return append(h.ToBytes(framecounter), []byte(payload)...)
}
