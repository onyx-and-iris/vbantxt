package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/onyx-and-iris/vbantxt"
	log "github.com/sirupsen/logrus"
)

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}

func main() {
	var (
		host       string
		port       int
		streamname string
		loglevel   int
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

	flag.IntVar(&bps, "bps", 0, "vban bps")
	flag.IntVar(&bps, "b", 0, "vban bps (shorthand)")
	flag.IntVar(&channel, "channel", 0, "vban channel")
	flag.IntVar(&channel, "c", 0, "vban channel (shorthand)")
	flag.IntVar(&ratelimit, "ratelimit", 20, "request ratelimit in milliseconds")
	flag.IntVar(&ratelimit, "r", 20, "request ratelimit in milliseconds (shorthand)")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		exit(err)
	}
	defaultConfigPath := filepath.Join(homeDir, ".config", "vbantxt", "config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "config path")
	flag.StringVar(&configPath, "C", defaultConfigPath, "config path (shorthand)")
	flag.IntVar(&loglevel, "loglevel", int(log.WarnLevel), "log level")
	flag.IntVar(&loglevel, "l", int(log.WarnLevel), "log level (shorthand)")
	flag.Parse()

	if slices.Contains(log.AllLevels, log.Level(loglevel)) {
		log.SetLevel(log.Level(loglevel))
	}

	if !flagsPassed([]string{"host", "h", "port", "p", "streamname", "s"}) {
		config, err := loadConfig(configPath)
		if err != nil {
			exit(err)
		}
		host = config.Host
		port = config.Port
		streamname = config.Streamname
	}

	client, err := createClient(host, port, streamname, bps, channel, ratelimit)
	if err != nil {
		exit(err)
	}
	defer client.Close()

	for _, arg := range flag.Args() {
		err := client.Send(arg)
		if err != nil {
			log.Error(err)
		}
	}
}

func createClient(host string, port int, streamname string, bps, channel, ratelimit int) (*vbantxt.VbanTxt, error) {
	client, err := vbantxt.New(
		host,
		port,
		streamname,
		vbantxt.WithBPSOpt(bps),
		vbantxt.WithChannel(channel),
		vbantxt.WithRateLimit(time.Duration(ratelimit)*time.Millisecond))
	if err != nil {
		return nil, err
	}
	return client, err
}
