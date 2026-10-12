package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"

	"github.com/onyx-and-iris/vbantxt"
	"github.com/onyx-and-iris/vbantxt/internal/logging"
)

type appContext struct {
	Ctx    context.Context
	Client *vbantxt.Client
	Out    io.Writer
}

type CLIConfig struct {
	Config   string `flag:"" default:"vbantxt.toml" help:"Path to the CLI configuration file." env:"VBANTXT_CLI_CONFIG_PATH"`
	LogLevel string `flag:"" default:"info"         help:"Log level for the CLI."              env:"VBANTXT_CLI_LOGLEVEL"    enum:"trace,debug,info,warn,error,fatal"`
}

type CLI struct {
	CLIConfig  CLIConfig          `embed:"" help:"CLI configuration"  group:"CLI"`
	VBANConfig vbantxt.VBANConfig `embed:"" help:"VBAN configuration" group:"VBAN"`

	Raw RawCmd `cmd:"" help:"Send raw VBAN-TXT payload."`

	Strip StripCmdGroup `help:"Control the strips." group:"Strip" cmd:""`
	Bus   BusCmdGroup   `help:"Control the buses."  group:"Bus"   cmd:""`
}

type RawCmd struct {
	Payload string `arg:"" help:"Raw VBAN-TXT payload"`
}

func (cmd *RawCmd) Run(ctx *appContext) error {
	return ctx.Client.SendTXT(ctx.Ctx, cmd.Payload)
}

func main() {
	var cli CLI

	initialParser, _ := buildInitialParser(&cli)
	configFromCLI, _ := extractConfigPath(initialParser, &cli)

	ctx := kong.Parse(
		&cli,
		kong.Configuration(kongtoml.Loader, configFromCLI, "/etc/vbantxt/app.toml"),
	)

	ctx.FatalIfErrorf(run(ctx, cli.CLIConfig, cli.VBANConfig))
}

func buildInitialParser(pre *CLI) (*kong.Kong, error) {
	return kong.New(pre,
		kong.Name("vbantxtctl"),
		kong.Description("VBAN-TXT control CLI"),
		kong.Help(func(options kong.HelpOptions, ctx *kong.Context) error {
			return nil
		}),
		kong.Exit(func(int) {}),
	)
}

func extractConfigPath(parser *kong.Kong, pre *CLI) (string, error) {
	if _, err := parser.Parse(os.Args[1:]); err != nil {
		return pre.CLIConfig.Config, fmt.Errorf("failed to parse CLI arguments: %w", err)
	}
	return pre.CLIConfig.Config, nil
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

	if err := client.Start(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
