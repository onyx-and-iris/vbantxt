```console
Usage: vbantxt <payload> ... [flags]

Arguments:
  <payload> ...    One or more raw VBAN-TXT payloads.

Flags:
  -h, --help    Show context-sensitive help.

CLI
  --config="vbantxt.toml"    Path to the CLI configuration file ($VBANTXT_CLI_CONFIG_PATH).
  --log-level="info"         Log level for the CLI ($VBANTXT_CLI_LOG_LEVEL).

VBAN
  --host="localhost"         ($VBANTXT_CLI_HOST)
  --port=6980                ($VBANTXT_CLI_PORT)
  --streamname="Command1"    ($VBANTXT_CLI_STREAMNAME)
  --bps=256000               ($VBANTXT_CLI_BPS)
  --channel=0                ($VBANTXT_CLI_CHANNEL)
  --rate-limit=20ms          ($VBANTXT_CLI_RATELIMIT)

```
