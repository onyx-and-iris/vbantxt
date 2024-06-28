# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Before any major/minor/patch bump all unit tests will be run to verify they pass.

## [Unreleased]

- [x]

# [0.1.0] - 2024-06-28

### Added

- Matrix and Logging sections to README.

### Changed

- `host` flag now defaults to "localhost". Useful if sending VBAN-Text to Matrix
- `loglevel` flag now expects values that correspond to the logrus package loglevels (0 up to 6). See README.
- Config values are only applied if the corresponding flag was not passed on the command line.

# [0.0.1] - 2022-09-23

### Added

- Initial release, package implements VBAN PROTOCOL TXT with a basic CLI for configuring options.
- Ability to load configuration settings from a config.toml.
