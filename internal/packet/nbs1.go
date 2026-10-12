package packet

import "fmt"

// VbanRTPacketNBS1 represents the body of an RTPacket with ident:1.
type VbanRTPacketNBS1 struct {
	Hdr   VbanRTPacketHeader
	Strip [16]VbanVMParamStrip
}

func (p *VbanRTPacketNBS1) NBS() NBS                    { return p.Hdr.NBSValue }
func (p *VbanRTPacketNBS1) Header() *VbanRTPacketHeader { return &p.Hdr }

func (p *VbanRTPacketNBS1) Equals(other *VbanRTPacketNBS1) bool {
	// NBS1 packets are always fresh; never stale.
	// Always treat them as dirty.
	return false
}

func ParseVbanRTPacketNBS1(nbs NBS, data []byte) (*VbanRTPacketNBS1, error) {
	hdr := parseHeader(nbs, data)

	const base = 44
	if len(data) < base {
		return nil, fmt.Errorf("NBS1 packet too short: %d", len(data))
	}

	p := &VbanRTPacketNBS1{
		Hdr:   hdr,
		Strip: [16]VbanVMParamStrip{},
	}

	// How many full strips are actually present in the packet?
	body := len(data) - base
	stripCount := body / VMPARAMSTRIP_SIZE
	if stripCount > 16 {
		stripCount = 16
	}

	// Parse only the strips that exist
	for i := 0; i < stripCount; i++ {
		start := base + i*VMPARAMSTRIP_SIZE
		end := start + VMPARAMSTRIP_SIZE

		if end > len(data) {
			// Should never happen, but guard anyway
			break
		}

		p.Strip[i] = ParseVbanVMParamStrip(data[start:end])
	}

	// Remaining strips stay zero-valued (same behavior as Python)
	return p, nil
}
