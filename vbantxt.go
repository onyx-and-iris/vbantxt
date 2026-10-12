package vbantxt

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/onyx-and-iris/vbantxt/internal/logging"
	"github.com/onyx-and-iris/vbantxt/packet"
)

// Client represents a VBAN TXT client that communicates with a VBAN server over UDP.
// It manages subscription, reading, and state tracking for NBS0 and NBS1 packets.
type Client struct {
	addr *net.UDPAddr
	conn *net.UDPConn

	hdr *packet.VbanRTRequestHeader

	rtChan chan packet.RTPacket

	ctx    context.Context
	cancel context.CancelFunc

	frameMu      sync.Mutex
	frameCounter uint32

	rateLimit time.Duration
	lastSend  time.Time
	sendMu    sync.Mutex

	lastNBS0 *packet.VbanRTPacketNBS0
	lastNBS1 *packet.VbanRTPacketNBS1
}

// New creates a new Client instance with the specified host, port, and stream name.
// It applies any provided options and returns the initialised client or an error if creation fails.
func New(host string, port int, streamname string, options ...Option) (*Client, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	hdr, err := packet.NewVbanRTRequestHeader(streamname)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	c := &Client{
		addr:         addr,
		conn:         conn,
		hdr:          hdr,
		rtChan:       make(chan packet.RTPacket, 64),
		ctx:          ctx,
		cancel:       cancel,
		frameCounter: 0,
		rateLimit:    time.Duration(20) * time.Millisecond,
		lastSend:     time.Time{},
	}

	for _, opt := range options {
		opt(c)
	}

	return c, nil
}

// Context returns the client's managed context.
func (c *Client) Context() context.Context {
	return c.ctx
}

// Start initialises the client and starts the subscriber and reader goroutines.
// It returns an error if initialization fails.
func (c *Client) Start(ctx context.Context) error {
	// tie external ctx to internal
	c.ctx, c.cancel = context.WithCancel(ctx)

	go c.runSubscriber()
	go c.runReader()

	return nil
}

func (c *Client) nextFrameCounter() uint32 {
	c.frameMu.Lock()
	defer c.frameMu.Unlock()

	c.frameCounter = (c.frameCounter + 1) & 0xFFFFFFFF
	return c.frameCounter
}

// runSubscriber continuously sends subscription packets for NBS0 and NBS1 to the server.
// It runs in its own goroutine and respects the client's context for cancellation.
func (c *Client) runSubscriber() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return

		default:
			for _, nbs := range []packet.NBS{packet.NBSZero, packet.NBSOne} {
				frame := c.nextFrameCounter()
				sub := packet.NewVbanRTSubscribeHeader(nbs, "Register-RTP", 15)
				pkt := sub.ToBytes(frame)
				logging.Tracef("Sending subscription packet for NBS %d, frame %d", nbs, frame)
				c.conn.Write(pkt)
			}
		}

		<-ticker.C
	}
}

// runReader continuously reads from the UDP connection, parses RT packets,
// and sends them to the rtChan if they are considered "dirty" (state is stale).
func (c *Client) runReader() {
	buf := make([]byte, 2048)

	for {
		select {
		case <-c.ctx.Done():
			close(c.rtChan)
			return

		default:
			n, err := c.conn.Read(buf)
			if err != nil || n < packet.HeaderSize {
				continue
			}

			// Validate VBAN header
			if _, err := packet.ParseVbanRTResponseHeader(buf[:packet.HeaderSize]); err != nil {
				continue
			}

			// Parse full packet
			pkt, err := packet.ParseRTPacket(buf[:n])
			if err != nil {
				logging.Errorf("Failed to parse RT packet: %v", err)
				continue
			}

			// Dirty filtering logic
			switch p := pkt.(type) {

			case *packet.VbanRTPacketNBS0:
				logging.Tracef("Received NBS0 packet")

				// Parameter dirty?
				pdirty := c.lastNBS0 == nil || !p.Equals(c.lastNBS0)
				logging.Tracef("Parameter dirty: %v", pdirty)

				// Levels dirty?
				ldirty := c.lastNBS0 == nil || !p.LevelsEqual(c.lastNBS0)
				logging.Tracef("Levels dirty: %v", ldirty)

				if pdirty || ldirty {
					c.lastNBS0 = p
					select {
					case c.rtChan <- p:
					case <-c.ctx.Done():
						close(c.rtChan)
						return
					}
				}

			case *packet.VbanRTPacketNBS1:
				logging.Tracef("Received NBS1 packet")
				// NBS1 only contains parameters, so only pdirty
				pdirty := c.lastNBS1 == nil || !p.Equals(c.lastNBS1)

				if pdirty {
					c.lastNBS1 = p
					select {
					case c.rtChan <- p:
					case <-c.ctx.Done():
						close(c.rtChan)
						return
					}
				}
			}
		}
	}
}

// RT returns the stream of fresh, dirty-filtered RT packets.
//
// Use this if you need direct control over the raw RT packet stream.
//
// For most use cases, consider using WaitForRT() or other high-level helpers instead.
func (c *Client) RT() <-chan packet.RTPacket {
	return c.rtChan
}

// WaitForRT waits for the next RT packet of the specified NBS type.
// It returns the packet once it is received or an error if the context is cancelled.
func (c *Client) WaitForRT(ctx context.Context, wantNBS packet.NBS) (packet.RTPacket, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while waiting for RT packet")

		case pkt := <-c.rtChan:
			switch p := pkt.(type) {
			case *packet.VbanRTPacketNBS0:
				if wantNBS == packet.NBSZero {
					return p, nil
				}
			case *packet.VbanRTPacketNBS1:
				if wantNBS == packet.NBSOne {
					return p, nil
				}
			}
		}
	}
}

// WaitForAnyRT waits for the next RT packet of any NBS type.
// It returns the packet once it is received or an error if the context is cancelled.
func (c *Client) WaitForAnyRT(ctx context.Context) (packet.RTPacket, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while waiting for RT packet")
		case pkt := <-c.rtChan:
			return pkt, nil
		}
	}
}

// SendTXT sends a TXT packet with the specified payload.
// It enforces the client's rate limit and returns an error if sending fails.
func (c *Client) SendTXT(ctx context.Context, payload string) error {
	c.sendMu.Lock()

	// enforce rate limit
	if !c.lastSend.IsZero() {
		elapsed := time.Since(c.lastSend)
		if elapsed < c.rateLimit {
			time.Sleep(c.rateLimit - elapsed)
		}
	}

	c.lastSend = time.Now()
	c.sendMu.Unlock()

	frame := c.nextFrameCounter()
	pkt := c.hdr.EncodeWithPayload(frame, payload)

	_, err := c.conn.Write(pkt)
	return err
}

// Close closes the client, cancelling its context and closing the underlying connection.
// It returns an error if closing the connection fails.
func (c *Client) Close() error {
	c.cancel()
	return c.conn.Close()
}
