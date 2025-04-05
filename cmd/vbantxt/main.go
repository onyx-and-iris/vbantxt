package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/onyx-and-iris/vbantxt"
	log "github.com/sirupsen/logrus"
)

type opts struct {
	host       string
	port       int
	streamname string
	bps        int
	channel    int
	ratelimit  int
	configPath string
	loglevel   string
}

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	var (
		host       string
		port       int
		streamname string
		loglevel   string
		configPath string
		bps        int
		channel    int
		ratelimit  int
	)

	flag.StringVar(&host, "host", "localhost", "vban host")
	flag.StringVar(&host, "h", "localhost", "vban host (shorthand)")
	flag.IntVar(&port, "port", 6980, "vban server port")
	flag.IntVar(&port, "p", 6980, "vban server port (shorthand)")
	flag.StringVar(&streamname, "streamname", "Command1", "stream name for text requests")
	flag.StringVar(&streamname, "s", "Command1", "stream name for text requests (shorthand)")

	flag.IntVar(&bps, "bps", 256000, "vban bps")
	flag.IntVar(&bps, "b", 256000, "vban bps (shorthand)")
	flag.IntVar(&channel, "channel", 0, "vban channel")
	flag.IntVar(&channel, "c", 0, "vban channel (shorthand)")
	flag.IntVar(&ratelimit, "ratelimit", 20, "request ratelimit in milliseconds")
	flag.IntVar(&ratelimit, "r", 20, "request ratelimit in milliseconds (shorthand)")

	configDir, err := os.UserConfigDir()
	if err != nil {
		exit(err)
	}
	defaultConfigPath := filepath.Join(configDir, "vbantxt", "config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "config path")
	flag.StringVar(&configPath, "C", defaultConfigPath, "config path (shorthand)")
	flag.StringVar(&loglevel, "loglevel", "warn", "log level")
	flag.StringVar(&loglevel, "l", "warn", "log level (shorthand)")

	flag.Parse()

	level, err := log.ParseLevel(loglevel)
	if err != nil {
		exit(fmt.Errorf("invalid log level: %s", loglevel))
	}
	log.SetLevel(level)

	o := opts{
		host:       host,
		port:       port,
		streamname: streamname,
		bps:        bps,
		channel:    channel,
		ratelimit:  ratelimit,
		configPath: configPath,
		loglevel:   loglevel,
	}

	// Load the config only if the host, port, and streamname flags are not provided.
	// This allows the user to override the config values with command line flags.
	if !flagsPassed([]string{"host", "h", "port", "p", "streamname", "s"}) {
		config, err := loadConfig(configPath)
		if err != nil {
			exit(err)
		}

		o.host = config.Host
		o.port = config.Port
		o.streamname = config.Streamname
	}
	log.Debugf("opts: %+v", o)

	client, closer, err := createClient(o)
	if err != nil {
		exit(err)
	}
	defer closer()

	for _, arg := range flag.Args() {
		err := client.Send(arg)
		if err != nil {
			log.Error(err)
		}
	}
}

// createClient creates a new vban client with the provided options.
func createClient(o opts) (*vbantxt.VbanTxt, func(), error) {
	client, err := vbantxt.New(
		o.host,
		o.port,
		o.streamname,
		vbantxt.WithBPSOpt(o.bps),
		vbantxt.WithChannel(o.channel),
		vbantxt.WithRateLimit(time.Duration(o.ratelimit)*time.Millisecond))
	if err != nil {
		return nil, nil, err
	}

	closer := func() {
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}

	return client, closer, err
}
