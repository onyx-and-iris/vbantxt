package packet

import (
	"encoding/binary"
)

type VbanRTSubscribeHeader struct {
	NBS     NBS
	Name    string
	Timeout uint8
}

func NewVbanRTSubscribeHeader(nbs NBS, name string, timeout uint8) *VbanRTSubscribeHeader {
	return &VbanRTSubscribeHeader{
		NBS:     nbs,
		Name:    name,
		Timeout: timeout,
	}
}

func (h *VbanRTSubscribeHeader) Streamname() [StreamnameMaxLength]byte {
	var out [StreamnameMaxLength]byte
	b := []byte(h.Name)
	if len(b) > StreamnameMaxLength {
		b = b[:StreamnameMaxLength]
	}
	copy(out[:], b)
	return out
}

// ToBytes serializes the subscription header into a byte slice suitable for sending over the network.
func (h *VbanRTSubscribeHeader) ToBytes(framecounter uint32) []byte {
	buf := make([]byte, 4+4+16+4)

	copy(buf[0:4], []byte("VBAN"))

	buf[4] = uint8(SubProtocolService)
	buf[5] = uint8(h.NBS & 0xFF)
	buf[6] = uint8(ServiceRTPacketRegister)
	buf[7] = h.Timeout

	sn := h.Streamname()
	copy(buf[8:24], sn[:])

	// framecounter (uint32 LE)
	binary.LittleEndian.PutUint32(buf[24:28], framecounter)

	return buf
}
