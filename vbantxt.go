package vbantxt

import (
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/log"
)

// VbanTxt is used to send VBAN-TXT requests to a distant Voicemeeter/Matrix.
type VbanTxt struct {
	conn      io.WriteCloser
	packet    packet
	ratelimit time.Duration
	lastSend  time.Time
}

// New constructs a fully formed VbanTxt instance. This is the package's entry point.
// It sets default values for it's fields and then runs the option functions.
func New(host string, port int, streamname string, options ...Option) (*VbanTxt, error) {
	conn, err := newUDPConn(host, port)
	if err != nil {
		return nil, fmt.Errorf("error creating UDP client for (%s:%d): %w", host, port, err)
	}

	packet, err := newPacket(streamname)
	if err != nil {
		return nil, fmt.Errorf("error creating packet: %w", err)
	}

	vt := &VbanTxt{
		conn:      conn,
		packet:    packet,
		ratelimit: time.Duration(20) * time.Millisecond,
	}

	for _, o := range options {
		o(vt)
	}

	return vt, nil
}

// Send is responsible for firing each VBAN-TXT request.
// It enforces rate limiting by waiting only when necessary.
func (vt *VbanTxt) Send(cmd string) error {
	if elapsed := time.Since(vt.lastSend); elapsed < vt.ratelimit {
		log.Debugf(
			"Rate limit in effect. Waiting for %v before sending next command.",
			vt.ratelimit-elapsed,
		)
		time.Sleep(vt.ratelimit - elapsed)
	}

	vt.lastSend = time.Now()

	_, err := vt.conn.Write(append(vt.packet.header(), cmd...))
	if err != nil {
		return fmt.Errorf("error sending command (%s): %w", cmd, err)
	}

	vt.packet.bumpFrameCounter()

	return nil
}

// Close is responsible for closing the UDP Client connection.
func (vt *VbanTxt) Close() error {
	err := vt.conn.Close()
	if err != nil {
		return fmt.Errorf("error attempting to close UDP Client: %w", err)
	}
	return nil
}
