package vbantxt

import (
	"encoding/binary"

	log "github.com/sirupsen/logrus"
)

const (
	vbanProtocolTxt = 0x40
	streamNameSz    = 16
	headerSz        = 4 + 1 + 1 + 1 + 1 + 16 + 4
)

var BpsOpts = []int{0, 110, 150, 300, 600, 1200, 2400, 4800, 9600, 14400, 19200, 31250,
	38400, 57600, 115200, 128000, 230400, 250000, 256000, 460800, 921600,
	1000000, 1500000, 2000000, 3000000}

type packet struct {
	name         string
	bpsIndex     int
	channel      int
	framecounter []byte
}

// newPacket returns a packet struct with default values, framecounter at 0.
func newPacket(streamname string) packet {
	return packet{
		name:         streamname,
		bpsIndex:     0,
		channel:      0,
		framecounter: make([]byte, 4),
	}
}

// sr defines the samplerate for the request
func (p *packet) sr() byte {
	return byte(vbanProtocolTxt + p.bpsIndex)
}

// nbc defines the channel of the request
func (p *packet) nbc() byte {
	return byte(p.channel)
}

// streamname defines the stream name of the text request
func (p *packet) streamname() []byte {
	b := make([]byte, streamNameSz)
	copy(b, p.name)
	return b
}

// header returns a fully formed packet header
func (p *packet) header() []byte {
	h := make([]byte, 0, headerSz)
	h = append(h, []byte("VBAN")...)
	h = append(h, p.sr())
	h = append(h, byte(0))
	h = append(h, p.nbc())
	h = append(h, byte(0x10))
	h = append(h, p.streamname()...)
	h = append(h, p.framecounter...)
	return h
}

// bumpFrameCounter increments the frame counter by 1
func (p *packet) bumpFrameCounter() {
	x := binary.LittleEndian.Uint32(p.framecounter)
	binary.LittleEndian.PutUint32(p.framecounter, x+1)

	log.Tracef("framecounter: %d", x)
}
