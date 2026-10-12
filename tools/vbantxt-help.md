```console
Usage: vbantxt <payload> ... [flags]

A VBAN-TXT command-line interface

Arguments:
  <payload> ...    One or more raw VBAN-TXT payloads.

Flags:
  -h, --help       Show context-sensitive help.
      --man        Print man page.
  -v, --version    Print vbantxt version information and quit

CLI
  -C, --config="vbantxt.toml"    Path to the CLI configuration file ($VBANTXT_CLI_CONFIG_PATH).
      --log-level="warn"         Log level for the CLI ($VBANTXT_CLI_LOG_LEVEL).

VBAN
  --host="localhost"         ($VBANTXT_CLI_HOST)
  --port=6980                ($VBANTXT_CLI_PORT)
  --streamname="Command1"    ($VBANTXT_CLI_STREAMNAME)
  --bps=256000               ($VBANTXT_CLI_BPS)
  --channel=0                ($VBANTXT_CLI_CHANNEL)
  --rate-limit=20ms          ($VBANTXT_CLI_RATELIMIT)

```
