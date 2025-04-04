# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Before any major/minor/patch bump all unit tests will be run to verify they pass.

## [Unreleased]

-   [x]

# [0.4.0] - 2024-04-04

### Changed

-   `log-level` flag is now of type string. It accepts any one of trace, debug, info, warn, error, fatal or panic.
    -   It defaults to warn.

# [0.3.1]

### Fixed

-   The CLI now uses `os.UserConfigDir()` to load the default *config.toml*, which should respect `$XDG_CONFIG_HOME`. See [UserConfigDir](https://pkg.go.dev/os#UserConfigDir)

# [0.2.1] - 2024-11-07

### Fixed

-   {packet}.header() now uses a reusable buffer.

# [0.2.0] - 2024-10-27

### Added

-   `config` flag (shorthand `C`), you may now specify a custom config directory. It defaults to `home directory / .config / vbantxt_cli /`.
    -   please note, the default directory has changed from v0.1.0
-   Functional options `WithRateLimit` and `WithBPSOpt` and `WithChannel` added. Use them to configure the vbantxt client. See the [included vbantxt cli][vbantxt-cli] for an example of usage.

### Changed

-   Behaviour change: if any one of `"host", "h", "port", "p", "streamname", "s"` flags are passed then the config file will be ignored.
-   `delay` flag changed to `ratelimit` (shorthand `r`). It defaults to 20ms.

# [0.1.0] - 2024-06-28

### Added

-   Matrix and Logging sections to README.

### Changed

-   `host` flag now defaults to "localhost". Useful if sending VBAN-Text to Matrix
-   `loglevel` flag now expects values that correspond to the logrus package loglevels (0 up to 6). See README.
-   Config values are only applied if the corresponding flag was not passed on the command line.

# [0.0.1] - 2022-09-23

### Added

-   Initial release, package implements VBAN PROTOCOL TXT with a basic CLI for configuring options.
-   Ability to load configuration settings from a config.toml.

[vbantxt-cli]: https://github.com/onyx-and-iris/vbantxt/blob/main/cmd/vbantxt/main.go
