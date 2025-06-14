package vbantxt

import (
	"fmt"
	"net"

	"github.com/charmbracelet/log"
)

// udpConn represents the UDP client.
type udpConn struct {
	conn *net.UDPConn
}

// newUDPConn returns a UDP client.
func newUDPConn(host string, port int) (udpConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return udpConn{}, err
	}
	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return udpConn{}, err
	}
	log.Infof("Outgoing address %s", conn.RemoteAddr())

	return udpConn{conn: conn}, nil
}

// Write implements the io.WriteCloser interface.
func (u udpConn) Write(buf []byte) (int, error) {
	n, err := u.conn.Write(buf)
	if err != nil {
		return 0, err
	}
	log.Debugf("Sending '%s' to: %s", string(buf), u.conn.RemoteAddr())

	return n, nil
}

// Close implements the io.WriteCloser interface.
func (u udpConn) Close() error {
	err := u.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
