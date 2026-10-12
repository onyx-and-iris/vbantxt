package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	mangokong "github.com/alecthomas/mango-kong"
	"github.com/muesli/roff"

	"github.com/onyx-and-iris/vbantxt"
	"github.com/onyx-and-iris/vbantxt/internal/logging"
)

var version string // Version of the CLI, set at build time.

// VersionFlag is a custom flag type that prints the version and exits.
type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }  // nolint: revive
func (v VersionFlag) IsBool() bool                       { return true } // nolint: revive
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error { // nolint: revive, unparam
	if app.Model.Name != "vbantxt" {
		return nil
	}

	fmt.Fprintf(app.Stdout, "vbantxt version: %s\n", vars["version"])
	app.Exit(0)
	return nil
}

type ManFlag mangokong.ManFlag

func (m ManFlag) Decode(_ *kong.DecodeContext) error { return nil }  // nolint: revive
func (m ManFlag) IsBool() bool                       { return true } // nolint: revive
func (m ManFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error { // nolint: revive, unparam
	if app.Model.Name != "vbantxt" {
		return nil
	}

	man := mangokong.NewManPage(1, app.Model)
	fmt.Fprint(app.Stdout, man.Build(roff.NewDocument()))
	app.Exit(0)
	return nil
}

type appContext struct {
	Ctx    context.Context
	Client *vbantxt.Client
	Out    io.Writer
}

type CLIConfig struct {
	Config   string `flag:"" default:"vbantxt.toml" help:"Path to the CLI configuration file." env:"VBANTXT_CLI_CONFIG_PATH" short:"C"`
	LogLevel string `flag:"" default:"warn"         help:"Log level for the CLI."              env:"VBANTXT_CLI_LOG_LEVEL"             enum:"trace,debug,info,warn,error,fatal"`
}

type CLI struct {
	CLIConfig  CLIConfig          `embed:"" help:"CLI configuration"  group:"CLI"`
	VBANConfig vbantxt.VBANConfig `embed:"" help:"VBAN configuration" group:"VBAN"`

	Man     ManFlag     `help:"Print man page."`
	Version VersionFlag `help:"Print vbantxt version information and quit" name:"version" short:"v"`

	RawCmd `cmd:"" help:"Send raw VBAN-TXT payload(s)."`
}

type RawCmd struct {
	Payload []string `arg:"" help:"One or more raw VBAN-TXT payloads."`
}

func (cmd *RawCmd) Run(ctx *appContext) error {
	for _, cmd := range cmd.Payload {
		if err := ctx.Client.SendTXT(ctx.Ctx, cmd); err != nil {
			logging.Errorf("Failed to send VBAN-TXT payload: %v", err)
		}
	}
	return nil
}

func main() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get user config directory: %v\n", err)
		os.Exit(1)
	}
	defaultConfigPath := filepath.Join(configDir, "vbantxt", "config.toml")

	var cli CLI

	initialParser, _ := buildInitialParser(&cli)
	configFromCLI := extractConfigPath(initialParser, &cli)

	ctx := kong.Parse(
		&cli,
		kong.Name("vbantxt"),
		kong.Description("A VBAN-TXT command-line interface"),
		kong.UsageOnError(),
		kong.Vars{
			"version": func() string {
				if version != "" {
					return version
				}

				info, ok := debug.ReadBuildInfo()
				if !ok {
					return "(unable to read version)"
				}
				return strings.Split(info.Main.Version, "-")[0]
			}(),
		},
		kong.Configuration(kongtoml.Loader, configFromCLI, defaultConfigPath),
	)

	ctx.FatalIfErrorf(run(ctx, cli.CLIConfig, cli.VBANConfig))
}

func buildInitialParser(pre *CLI) (*kong.Kong, error) {
	return kong.New(pre,
		kong.Name("pre"),
		kong.Description("First pass parser"),
		kong.Help(func(options kong.HelpOptions, ctx *kong.Context) error {
			return nil
		}),
		kong.Exit(func(int) {}),
	)
}

func extractConfigPath(parser *kong.Kong, pre *CLI) string {
	if _, err := parser.Parse(os.Args[1:]); err != nil {
		return "vbantxt.toml"
	}
	return pre.CLIConfig.Config
}

func run(ctx *kong.Context, cliConfig CLIConfig, vbanConfig vbantxt.VBANConfig) error {
	loglevel, err := logging.ParseLevel(cliConfig.LogLevel)
	if err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}
	logging.SetLevel(loglevel)

	client, err := initialise(vbanConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx.Bind(&appContext{
		Ctx:    context.Background(),
		Client: client,
		Out:    ctx.Stdout,
	})

	return ctx.Run()
}

func initialise(cfg vbantxt.VBANConfig) (*vbantxt.Client, error) {
	client, err := vbantxt.New(
		cfg.Host,
		cfg.Port,
		cfg.Streamname,
		vbantxt.WithRateLimit(cfg.RateLimit),
		vbantxt.WithBPSOpt(cfg.BPS),
		vbantxt.WithChannel(cfg.Channel),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
