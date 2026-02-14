package vbantxt

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/charmbracelet/log"
)

const (
	vbanProtocolTxt = 0x40
	streamNameSz    = 16
	headerSz        = 4 + 1 + 1 + 1 + 1 + 16 + 4
)

// BpsOpts defines the available baud rate options.
var BpsOpts = []int{
	0, 110, 150, 300, 600, 1200, 2400, 4800, 9600, 14400, 19200, 31250,
	38400, 57600, 115200, 128000, 230400, 250000, 256000, 460800, 921600,
	1000000, 1500000, 2000000, 3000000,
}

type packet struct {
	streamname   [streamNameSz]byte
	bpsIndex     uint8
	channel      uint8
	framecounter uint32
	hbuf         *bytes.Buffer
}

// newPacket creates a new packet with the given stream name and default values for other fields.
// It validates the stream name length and ensures the default baud rate is present in BpsOpts.
func newPacket(streamname string) (packet, error) {
	if len(streamname) > streamNameSz {
		return packet{}, fmt.Errorf(
			"streamname too long: %d chars, max %d",
			len(streamname),
			streamNameSz,
		)
	}

	var streamnameBuf [streamNameSz]byte
	copy(streamnameBuf[:], streamname)

	bpsIndex := indexOf(BpsOpts, 256000)
	if bpsIndex == -1 {
		return packet{}, fmt.Errorf("default baud rate 256000 not found in BpsOpts")
	}

	return packet{
		streamname:   streamnameBuf,
		bpsIndex:     uint8(bpsIndex),
		channel:      0,
		framecounter: 0,
		hbuf:         bytes.NewBuffer(make([]byte, 0, headerSz)),
	}, nil
}

// sr defines the samplerate for the request.
func (p *packet) sr() byte {
	return byte(vbanProtocolTxt + p.bpsIndex)
}

// nbc defines the channel of the request.
func (p *packet) nbc() byte {
	return byte(p.channel)
}

// header returns a fully formed packet header.
func (p *packet) header() []byte {
	p.hbuf.Reset()
	p.hbuf.WriteString("VBAN")
	p.hbuf.WriteByte(p.sr())
	p.hbuf.WriteByte(byte(0))
	p.hbuf.WriteByte(p.nbc())
	p.hbuf.WriteByte(byte(0x10))
	p.hbuf.Write(p.streamname[:])

	var frameBytes [4]byte
	binary.LittleEndian.PutUint32(frameBytes[:], p.framecounter)
	p.hbuf.Write(frameBytes[:])

	return p.hbuf.Bytes()
}

// bumpFrameCounter increments the frame counter by 1.
// The uint32 will safely wrap to 0 after reaching max value (4,294,967,295),
// which is expected behaviour for network protocol sequence numbers.
func (p *packet) bumpFrameCounter() {
	p.framecounter++

	log.Debugf("framecounter: %d", p.framecounter)
}
