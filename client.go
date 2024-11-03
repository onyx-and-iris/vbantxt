package vbantxt

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

// client represents the UDP client
type client struct {
	conn *net.UDPConn
}

// NewClient returns a UDP client
func newClient(host string, port int) (client, error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return client{}, err
	}
	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return client{}, err
	}
	log.Infof("Outgoing address %s", conn.RemoteAddr())

	return client{conn: conn}, nil
}

// Write implements the io.WriteCloser interface
func (c client) Write(buf []byte) (int, error) {
	n, err := c.conn.Write(buf)
	if err != nil {
		return 0, err
	}
	log.Debugf("Sending '%s' to: %s", string(buf), c.conn.RemoteAddr())

	return n, nil
}

// Close implements the io.WriteCloser interface
func (c client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
