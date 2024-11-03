package vbantxt

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Option is a functional option type that allows us to configure the VbanTxt.
type Option func(*VbanTxt)

// WithRateLimit is a functional option to set the ratelimit for requests
func WithRateLimit(ratelimit time.Duration) Option {
	return func(vt *VbanTxt) {
		vt.ratelimit = ratelimit
	}
}

// WithBPSOpt is a functional option to set the bps index for {VbanTx}.{Packet}.bpsIndex
func WithBPSOpt(bpsIndex int) Option {
	return func(vt *VbanTxt) {
		if bpsIndex < 0 || bpsIndex >= len(BpsOpts) {
			log.Warnf("invalid bpsIndex %d, defaulting to 0", bpsIndex)
			return
		}
		vt.packet.bpsIndex = bpsIndex
	}
}

// WithChannel is a functional option to set the bps index for {VbanTx}.{Packet}.channel
func WithChannel(channel int) Option {
	return func(vt *VbanTxt) {
		vt.packet.channel = channel
	}
}
