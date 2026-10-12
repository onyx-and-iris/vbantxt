package vbantxt

import (
	"time"

	"github.com/onyx-and-iris/vbantxt/internal/logging"
	"github.com/onyx-and-iris/vbantxt/internal/packet"
)

// Option is a functional option type that allows us to configure the VbanTxt.
type Option func(*Client)

// WithRateLimit is a functional option to set the ratelimit for requests.
func WithRateLimit(ratelimit time.Duration) Option {
	return func(c *Client) {
		c.rateLimit = ratelimit
	}
}

// WithBPSOpt is a functional option to set the bps index for {VbanTxt}.packet.
func WithBPSOpt(bps int) Option {
	return func(c *Client) {
		defaultBps := packet.BpsOpts[c.hdr.BPSIndex]

		bpsIndex, err := packet.BpsIndex(bps)
		if err != nil {
			logging.Warnf(
				"invalid bps value %d, expected one of %v, defaulting to %d",
				bps,
				packet.BpsOpts,
				defaultBps,
			)
			return
		}
		c.hdr.BPSIndex = uint8(bpsIndex)
	}
}

// WithChannel is a functional option to set the channel for {VbanTxt}.packet.
func WithChannel(channel int) Option {
	return func(c *Client) {
		if channel < 0 || channel > 255 {
			logging.Warnf("channel value %d out of range [0,255], keeping default %d", channel, c.hdr.Channel)
			return
		}
		c.hdr.Channel = uint8(channel)
	}
}
