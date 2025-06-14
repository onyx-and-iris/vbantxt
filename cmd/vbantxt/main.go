// Package main implements a command-line tool for sending text messages over VBAN using the vbantxt library.
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/onyx-and-iris/vbantxt"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
	"github.com/peterbourgon/ff/v4/fftoml"
)

var version string // Version will be set at build time

// Flags holds the command-line flags for the VBANTXT client.
type Flags struct {
	Host       string
	Port       int
	Streamname string
	Bps        int
	Channel    int
	Ratelimit  int
	ConfigPath string // Path to the configuration file
	Loglevel   string // Log level
	Version    bool   // Version flag
}

func exitOnError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	var flags Flags

	// VBAN specific flags
	fs := ff.NewFlagSet("vbantxt")
	fs.StringVar(&flags.Host, 'H', "host", "localhost", "VBAN host")
	fs.IntVar(&flags.Port, 'p', "port", 6980, "VBAN port")
	fs.StringVar(&flags.Streamname, 's', "streamname", "Command1", "VBAN stream name")
	fs.IntVar(&flags.Bps, 'b', "bps", 256000, "VBAN BPS")
	fs.IntVar(&flags.Channel, 'n', "channel", 0, "VBAN channel")
	fs.IntVar(&flags.Ratelimit, 'r', "ratelimit", 20, "VBAN rate limit (ms)")

	configDir, err := os.UserConfigDir()
	if err != nil {
		exitOnError(fmt.Errorf("failed to get user config directory: %w", err))
	}
	defaultConfigPath := filepath.Join(configDir, "vbantxt", "config.toml")

	// Configuration file and logging flags
	fs.StringVar(&flags.ConfigPath, 'C', "config", defaultConfigPath, "Path to the configuration file")
	fs.StringVar(&flags.Loglevel, 'l', "loglevel", "warn", "Log level (debug, info, warn, error, fatal, panic)")
	fs.BoolVar(&flags.Version, 'v', "version", "Show version information")

	err = ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("VBANTXT"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigAllowMissingFile(),
		ff.WithConfigFileParser(fftoml.Parser{Delimiter: "."}.Parse),
	)
	switch {
	case errors.Is(err, ff.ErrHelp):
		fmt.Fprintf(os.Stderr, "%s\n", ffhelp.Flags(fs, "vbantxt [flags] <vban commands>"))
		os.Exit(0)
	case err != nil:
		exitOnError(fmt.Errorf("failed to parse flags: %w", err))
	}

	if flags.Version {
		if version == "" {
			info, ok := debug.ReadBuildInfo()
			if !ok {
				exitOnError(errors.New("failed to read build info"))
			}
			version = strings.Split(info.Main.Version, "-")[0]
		}
		fmt.Printf("vbantxt version: %s\n", version)
		os.Exit(0)
	}

	level, err := log.ParseLevel(flags.Loglevel)
	if err != nil {
		exitOnError(fmt.Errorf("invalid log level: %s", flags.Loglevel))
	}
	log.SetLevel(level)

	log.Debugf("Loaded configuration: %+v", flags)

	client, closer, err := createClient(&flags)
	if err != nil {
		exitOnError(err)
	}
	defer closer()

	for _, arg := range fs.GetArgs() {
		err := client.Send(arg)
		if err != nil {
			log.Error(err)
		}
	}
}

// createClient creates a new vban client with the provided options.
func createClient(flags *Flags) (*vbantxt.VbanTxt, func(), error) {
	client, err := vbantxt.New(
		flags.Host,
		flags.Port,
		flags.Streamname,
		vbantxt.WithBPSOpt(flags.Bps),
		vbantxt.WithChannel(flags.Channel),
		vbantxt.WithRateLimit(time.Duration(flags.Ratelimit)*time.Millisecond))
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
