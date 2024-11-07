package vbantxt

import (
	"bytes"
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
	streamname   []byte
	bpsIndex     int
	channel      int
	framecounter []byte
	hbuf         bytes.Buffer
}

// newPacket returns a packet struct with default values, framecounter at 0.
func newPacket(streamname string) packet {
	streamnameBuf := make([]byte, streamNameSz)
	copy(streamnameBuf, streamname)

	return packet{
		streamname:   streamnameBuf,
		bpsIndex:     0,
		channel:      0,
		framecounter: make([]byte, 4),
		hbuf:         *bytes.NewBuffer(make([]byte, headerSz)),
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

// header returns a fully formed packet header
func (p *packet) header() []byte {
	p.hbuf.Reset()
	p.hbuf.WriteString("VBAN")
	p.hbuf.WriteByte(p.sr())
	p.hbuf.WriteByte(byte(0))
	p.hbuf.WriteByte(p.nbc())
	p.hbuf.WriteByte(byte(0x10))
	p.hbuf.Write(p.streamname)
	p.hbuf.Write(p.framecounter)
	return p.hbuf.Bytes()
}

// bumpFrameCounter increments the frame counter by 1
func (p *packet) bumpFrameCounter() {
	x := binary.LittleEndian.Uint32(p.framecounter)
	binary.LittleEndian.PutUint32(p.framecounter, x+1)

	log.Tracef("framecounter: %d", x)
}
