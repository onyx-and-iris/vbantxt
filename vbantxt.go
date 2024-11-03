package vbantxt

import (
	"fmt"
	"io"
	"time"
)

// VbanTxt is used to send VBAN-TXT requests to a distant Voicemeeter/Matrix.
type VbanTxt struct {
	client    io.WriteCloser
	packet    packet
	ratelimit time.Duration
}

// New constructs a fully formed VbanTxt instance. This is the package's entry point.
// It sets default values for it's fields and then runs the option functions.
func New(host string, port int, streamname string, options ...Option) (*VbanTxt, error) {
	client, err := newClient(host, port)
	if err != nil {
		return nil, fmt.Errorf("error creating UDP client for (%s:%d): %w", host, port, err)
	}

	vt := &VbanTxt{
		client:    client,
		packet:    newPacket(streamname),
		ratelimit: time.Duration(20) * time.Millisecond,
	}

	for _, o := range options {
		o(vt)
	}

	return vt, nil
}

// Send is resonsible for firing each VBAN-TXT request.
// It waits for {vt.ratelimit} time before returning.
func (vt VbanTxt) Send(cmd string) error {
	_, err := vt.client.Write(append(vt.packet.header(), []byte(cmd)...))
	if err != nil {
		return fmt.Errorf("error sending command (%s): %w", cmd, err)
	}

	vt.packet.bumpFrameCounter()

	time.Sleep(vt.ratelimit)

	return nil
}

// Close is responsible for closing the UDP Client connection
func (vt VbanTxt) Close() error {
	err := vt.client.Close()
	if err != nil {
		return fmt.Errorf("error attempting to close UDP Client: %w", err)
	}
	return nil
}
