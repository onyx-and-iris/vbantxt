package vbantxt

import (
	"time"

	"github.com/charmbracelet/log"
)

// Option is a functional option type that allows us to configure the VbanTxt.
type Option func(*VbanTxt)

// WithRateLimit is a functional option to set the ratelimit for requests.
func WithRateLimit(ratelimit time.Duration) Option {
	return func(vt *VbanTxt) {
		vt.ratelimit = ratelimit
	}
}

// WithBPSOpt is a functional option to set the bps index for {VbanTxt}.packet.
func WithBPSOpt(bps int) Option {
	return func(vt *VbanTxt) {
		defaultBps := BpsOpts[vt.packet.bpsIndex]

		bpsIndex := indexOf(BpsOpts, bps)
		if bpsIndex == -1 {
			log.Warnf(
				"invalid bps value %d, expected one of %v, defaulting to %d",
				bps,
				BpsOpts,
				defaultBps,
			)
			return
		}
		if bpsIndex > 255 {
			log.Warnf("bps index %d too large for uint8, defaulting to %d", bpsIndex, defaultBps)
			return
		}
		vt.packet.bpsIndex = uint8(bpsIndex)
	}
}

// WithChannel is a functional option to set the channel for {VbanTxt}.packet.
func WithChannel(channel int) Option {
	return func(vt *VbanTxt) {
		if channel < 0 || channel > 255 {
			log.Warnf(
				"channel value %d out of range [0,255], defaulting to %d",
				channel,
				vt.packet.channel,
			)
			return
		}
		vt.packet.channel = uint8(channel)
	}
}
