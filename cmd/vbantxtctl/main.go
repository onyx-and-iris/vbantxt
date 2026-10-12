package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"

	"github.com/onyx-and-iris/vbantxt"
	"github.com/onyx-and-iris/vbantxt/internal/logging"
)

type appContext struct {
	Ctx     context.Context
	Client  *vbantxt.Client
	Out     io.Writer
	Timeout time.Duration
}

type CLIConfig struct {
	Config   string        `flag:"" default:"vbantxt.toml" help:"Path to the CLI configuration file." env:"VBANTXT_CLI_CONFIG_PATH"`
	Timeout  time.Duration `flag:"" default:"2s"           help:"Timeout for RT operations."`
	LogLevel string        `flag:"" default:"info"         help:"Log level for the CLI."              env:"VBANTXT_CLI_LOG_LEVEL"   enum:"trace,debug,info,warn,error,fatal"`
}

type CLI struct {
	CLIConfig  CLIConfig          `embed:"" help:"CLI configuration"  group:"CLI"`
	VBANConfig vbantxt.VBANConfig `embed:"" help:"VBAN configuration" group:"VBAN"`

	Raw RawCmd `cmd:"" help:"Send raw VBAN-TXT payload."`

	Strip StripCmdGroup `help:"Control the strips." group:"Strip" cmd:""`
	Bus   BusCmdGroup   `help:"Control the buses."  group:"Bus"   cmd:""`
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
	var cli CLI

	initialParser, _ := buildInitialParser(&cli)
	configFromCLI := extractConfigPath(initialParser, &cli)

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

func extractConfigPath(parser *kong.Kong, pre *CLI) string {
	parser.Parse(os.Args[1:])
	return pre.CLIConfig.Config
}

func run(ctx *kong.Context, cliConfig CLIConfig, vbanConfig vbantxt.VBANConfig) error {
	loglevel, err := logging.ParseLevel(cliConfig.LogLevel)
	if err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}
	logging.SetLevel(loglevel)

	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	client, err := createClient(vbanConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.Start(rootCtx); err != nil {
		return err
	}

	ctx.Bind(&appContext{
		Ctx:     rootCtx,
		Client:  client,
		Out:     ctx.Stdout,
		Timeout: cliConfig.Timeout,
	})

	return ctx.Run()
}

func createClient(cfg vbantxt.VBANConfig) (*vbantxt.Client, error) {
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
