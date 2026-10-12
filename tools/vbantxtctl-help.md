```console
Usage: vbantxtctl <command> [flags]

Flags:
  -h, --help    Show context-sensitive help.

CLI
  --config="vbantxt.toml"    Path to the CLI configuration file ($VBANTXT_CLI_CONFIG_PATH).
  --timeout=2s               Timeout for RT operations.
  --log-level="info"         Log level for the CLI ($VBANTXT_CLI_LOG_LEVEL).

VBAN
  --host="localhost"         ($VBANTXT_CLI_HOST)
  --port=6980                ($VBANTXT_CLI_PORT)
  --streamname="Command1"    ($VBANTXT_CLI_STREAMNAME)
  --bps=256000               ($VBANTXT_CLI_BPS)
  --channel=0                ($VBANTXT_CLI_CHANNEL)
  --rate-limit=20ms          ($VBANTXT_CLI_RATELIMIT)

Commands:
  raw <payload> ... [flags]
    Send raw VBAN-TXT payload.

Strip
  strip <index> mono [<state>]
    Get or set the mono state of the strip.

  strip <index> solo [<state>]
    Get or set the solo state of the strip.

  strip <index> mute [<state>]
    Get or set the mute state of the strip.

  strip <index> gain [<state>]
    Get or set the gain of the strip.

  strip <index> label [<label>]
    Get or set the label of the strip.

  strip <index> a1 [<state>]
    Get or set the A1 state of the strip.

  strip <index> a2 [<state>]
    Get or set the A2 state of the strip.

  strip <index> a3 [<state>]
    Get or set the A3 state of the strip.

  strip <index> a4 [<state>]
    Get or set the A4 state of the strip.

  strip <index> a5 [<state>]
    Get or set the A5 state of the strip.

  strip <index> b1 [<state>]
    Get or set the B1 state of the strip.

  strip <index> b2 [<state>]
    Get or set the B2 state of the strip.

  strip <index> b3 [<state>]
    Get or set the B3 state of the strip.

  strip <index> eq on [<state>]
    Get or set strip EQ on/off.

  strip <index> eq channel <index> cell <index> on [<state>]
    Get/set EQ cell on/off.

  strip <index> eq channel <index> cell <index> type [<state>]
    Get/set EQ cell type.

  strip <index> eq channel <index> cell <index> freq [<state>]
    Get/set EQ cell frequency.

  strip <index> eq channel <index> cell <index> gain [<state>]
    Get/set EQ cell gain.

  strip <index> eq channel <index> cell <index> q [<state>]
    Get/set EQ cell Q factor.

  strip <index> comp knob [<state>]
    Get or set the knob settings of the compressor.

  strip <index> comp input-gain [<state>]
    Get or set the input gain of the compressor.

  strip <index> comp ratio [<state>]
    Get or set the ratio of the compressor.

  strip <index> comp threshold [<state>]
    Get or set the threshold of the compressor.

  strip <index> comp attack [<state>]
    Get or set the attack of the compressor.

  strip <index> comp release [<state>]
    Get or set the release of the compressor.

  strip <index> comp knee [<state>]
    Get or set the knee of the compressor.

  strip <index> comp output-gain [<state>]
    Get or set the output gain of the compressor.

  strip <index> comp makeup [<state>]
    Get or set the makeup gain of the compressor.

  strip <index> gate knob [<state>]
    Get or set the knob settings of the gate.

  strip <index> gate threshold [<state>]
    Get or set the threshold of the gate.

  strip <index> gate damping-max [<state>]
    Get or set the maximum damping of the gate.

  strip <index> gate bp-sidechain [<state>]
    Get or set the sidechain of the gate.

  strip <index> gate attack [<state>]
    Get or set the attack of the gate.

  strip <index> gate hold [<state>]
    Get or set the hold of the gate.

  strip <index> gate release [<state>]
    Get or set the release of the gate.

  strip <index> denoiser knob [<state>]
    Get or set the knob settings of the denoiser.

  strip <index> denoiser threshold [<state>]
    Get or set the threshold of the denoiser.

  strip <index> gainlayer <index> level [<level>]
    Control the level of the gain layer.

Bus
  bus <index> mono [<state>]
    Get or set the mono state of the bus.

  bus <index> mute [<state>]
    Get or set the mute state of the bus.

  bus <index> mode [<state>]
    Get or set the mode of the bus.

  bus <index> gain [<state>]
    Get or set the gain of the bus.

  bus <index> label [<state>]
    Get or set the label of the bus.

  bus <index> eq on [<state>]
    Get or set bus EQ on/off.

  bus <index> eq channel <index> cell <index> on [<state>]

  bus <index> eq channel <index> cell <index> type [flags]

  bus <index> eq channel <index> cell <index> freq [flags]

  bus <index> eq channel <index> cell <index> gain [flags]

  bus <index> eq channel <index> cell <index> q [flags]

Run "vbantxtctl <command> --help" for more information on a command.

```
