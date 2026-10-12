```console
Usage: vbantxtctl <command> [flags]

A VBAN-TXT command-line interface supporting RT packets

Flags:
  -h, --help       Show context-sensitive help.
      --man        Print man page.
  -v, --version    Print vbantxtctl version information and quit

CLI
  --config="vbantxt.toml"    Path to the CLI configuration file ($VBANTXT_CLI_CONFIG_PATH).
  --timeout=2s               Timeout for RT operations ($VBANTXT_CLI_TIMEOUT).
  --log-level="warn"         Log level for the CLI ($VBANTXT_CLI_LOG_LEVEL).

VBAN
  --host="localhost"         ($VBANTXT_CLI_HOST)
  --port=6980                ($VBANTXT_CLI_PORT)
  --streamname="Command1"    ($VBANTXT_CLI_STREAMNAME)
  --bps=256000               ($VBANTXT_CLI_BPS)
  --channel=0                ($VBANTXT_CLI_CHANNEL)
  --rate-limit=20ms          ($VBANTXT_CLI_RATELIMIT)

Commands:
  raw    Send raw VBAN-TXT payload.

Strip
  strip <index> mono                                    Get or set the mono state of the strip.
  strip <index> solo                                    Get or set the solo state of the strip.
  strip <index> mute                                    Get or set the mute state of the strip.
  strip <index> gain                                    Get or set the gain of the strip.
  strip <index> label                                   Get or set the label of the strip.
  strip <index> a1                                      Get or set the A1 state of the strip.
  strip <index> a2                                      Get or set the A2 state of the strip.
  strip <index> a3                                      Get or set the A3 state of the strip.
  strip <index> a4                                      Get or set the A4 state of the strip.
  strip <index> a5                                      Get or set the A5 state of the strip.
  strip <index> b1                                      Get or set the B1 state of the strip.
  strip <index> b2                                      Get or set the B2 state of the strip.
  strip <index> b3                                      Get or set the B3 state of the strip.
  strip <index> eq on                                   Get or set strip EQ on/off.
  strip <index> eq channel <index> cell <index> on      Get/set EQ cell on/off.
  strip <index> eq channel <index> cell <index> type    Get/set EQ cell type.
  strip <index> eq channel <index> cell <index> freq    Get/set EQ cell frequency.
  strip <index> eq channel <index> cell <index> gain    Get/set EQ cell gain.
  strip <index> eq channel <index> cell <index> q       Get/set EQ cell Q factor.
  strip <index> comp knob                               Get or set the knob settings of the compressor.
  strip <index> comp input-gain                         Get or set the input gain of the compressor.
  strip <index> comp ratio                              Get or set the ratio of the compressor.
  strip <index> comp threshold                          Get or set the threshold of the compressor.
  strip <index> comp attack                             Get or set the attack of the compressor.
  strip <index> comp release                            Get or set the release of the compressor.
  strip <index> comp knee                               Get or set the knee of the compressor.
  strip <index> comp output-gain                        Get or set the output gain of the compressor.
  strip <index> comp makeup                             Get or set the makeup gain of the compressor.
  strip <index> gate knob                               Get or set the knob settings of the gate.
  strip <index> gate threshold                          Get or set the threshold of the gate.
  strip <index> gate damping-max                        Get or set the maximum damping of the gate.
  strip <index> gate bp-sidechain                       Get or set the sidechain of the gate.
  strip <index> gate attack                             Get or set the attack of the gate.
  strip <index> gate hold                               Get or set the hold of the gate.
  strip <index> gate release                            Get or set the release of the gate.
  strip <index> denoiser knob                           Get or set the knob settings of the denoiser.
  strip <index> denoiser threshold                      Get or set the threshold of the denoiser.
  strip <index> gainlayer <index> level                 Control the level of the gain layer.

Bus
  bus <index> mono                                    Get or set the mono state of the bus.
  bus <index> mute                                    Get or set the mute state of the bus.
  bus <index> mode                                    Get or set the mode of the bus.
  bus <index> gain                                    Get or set the gain of the bus.
  bus <index> label                                   Get or set the label of the bus.
  bus <index> eq on                                   Get or set bus EQ on/off.
  bus <index> eq channel <index> cell <index> on
  bus <index> eq channel <index> cell <index> type
  bus <index> eq channel <index> cell <index> freq
  bus <index> eq channel <index> cell <index> gain
  bus <index> eq channel <index> cell <index> q

Run "vbantxtctl <command> --help" for more information on a command.

```
