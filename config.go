package vbantxt

import "time"

type VBANConfig struct {
	Host       string        `flag:"" default:"localhost" env:"VBANTXT_CLI_HOST"`
	Port       int           `flag:"" default:"6980"      env:"VBANTXT_CLI_PORT"`
	Streamname string        `flag:"" default:"Command1"  env:"VBANTXT_CLI_STREAMNAME"`
	BPS        int           `flag:"" default:"256000"    env:"VBANTXT_CLI_BPS"`
	Channel    int           `flag:"" default:"0"         env:"VBANTXT_CLI_CHANNEL"`
	RateLimit  time.Duration `flag:"" default:"20ms"      env:"VBANTXT_CLI_RATELIMIT"`
}
