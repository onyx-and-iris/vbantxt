![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)

# VBAN Sendtext

Send Voicemeeter/Matrix vban requests.

For an outline of past/future changes refer to: [CHANGELOG](CHANGELOG.md)

---

## Table of Contents

- [Installation](#installation)
- [VBANTXT Package](#vbantxt-package)
- [VBANTXT CLI](#vbantxt-cli)
- [License](#license)

## Requirements

-   [Voicemeeter](https://voicemeeter.com/) or [Matrix](https://vb-audio.com/Matrix/)
-   Go 1.18 or greater (a binary is available in [Releases](https://github.com/onyx-and-iris/vbantxt/releases))

## Tested against

-   Basic 1.1.2.2
-   Banana 2.1.2.2
-   Potato 3.1.2.2
-   Matrix 1.0.2.6

---

## `VBANTXT Package`

```console
go get github.com/onyx-and-iris/vbantxt
```

### SendTXT

```go
package main

import (
    "log"

    "github.com/onyx-and-iris/vbantxt"
)

func main() {
    client, err := vbantxt.New("vm.local", 6980, "onyx")
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }
    defer client.Close()

    if err := client.SendTXT("strip[0].mute=0"); err != nil {
        log.Fatalf("failed to send TXT command: %v", err)
    }
}
```

### SendTXT + RT

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/onyx-and-iris/vbantxt"
    "github.com/onyx-and-iris/vbantxt/packet"
)

func main() {
    client, err := vbantxt.New("vm.local", 6980, "onyx")
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }
    defer client.Close()

    // Root context (can be cancelled on Ctrl‑C in real applications)
    rootCtx := context.Background()

    // Start subscriber + reader goroutines
    if err := client.Start(rootCtx); err != nil {
        log.Fatalf("failed to start client: %v", err)
    }

    if err := client.SendTXT(rootCtx, "strip[0].mute=0"); err != nil {
        log.Fatalf("failed to send TXT command: %v", err)
    }

    // Operation context with timeout
    opCtx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
    defer cancel()

    // Wait for the next NBS0 packet
    pkt, err := client.WaitForRT(opCtx, packet.NBSZero)
    if err != nil {
        log.Fatalf("failed to receive RT packet: %v", err)
    }

    // Read back the mute state
    nbs0 := pkt.(*packet.VbanRTPacketNBS0)
    muted := nbs0.States().Strip[0].Mute()

    fmt.Printf("strip 0 mute: %v\n", muted)
}
```

---

## `VBANTXT CLI`

### Installation

```console
go install github.com/onyx-and-iris/vbantxt/cmd/vbantxt@latest
```

### Use

Simply pass your vban commands as command line arguments:

```console
vbantxt "strip[0].mute=1 strip[1].mono=1"
```

---

## `VBANTXTCTL CLI`

### Installation

```console
go install github.com/onyx-and-iris/vbantxt/cmd/vbantxtctl@latest
```

### Use

Instead of raw string requests vbantxtcli uses a command tree parser.

*set strip 0 mute*
```console
vbantxtctl strip 0 mute true
```

*read back value of strip 0 mute*
```console
vbantxtctl strip 0 mute
```

---

### Configuration

*vbantxt*

Please refer to [vbantxt-help](./tools/vbantxt-help.md)

*vbantxtctl*

Please refer to [vbantxtctl-help](./tools/vbantxtctl-help.md)


#### Environment Variables

All flags have corresponding environment variables, prefixed with `VBANTXT_CLI_`:

```bash
#!/usr/bin/env bash

export VBANTXT_CLI_HOST=localhost
export VBANTXT_CLI_PORT=6980
export VBANTXT_CLI_STREAMNAME=Command1
```

Flags will override environment variables.

#### TOML Config

By default the config loader will look for a config in:

-	$XDG_CONFIG_HOME / vbantxt / config.toml (see [os.UserConfigDir](https://pkg.go.dev/os#UserConfigDir))
	-	A custom config path may be passed with the --config/-C flag.

All flags have corresponding keys in the config file, for example:

```toml
host="gamepc.local"
port=6980
streamname="Command1"
```

---

## `Script files`

The vbantxt CLI accepts a single string request or an array of string requests. This means you can pass scripts stored in files.

For example, in Windows with Powershell you could:

```console
vbantxt $(Get-Content .\script.txt)
```

Or with Bash:

```console
xargs vbantxt < script.txt
```

to load commands from a file:

```
strip[0].mute=1;strip[0].mono=0
strip[1].mute=0;strip[1].mono=1
bus[3].eq.On=0
```

---

## `Matrix`

Sending commands to VB-Audio Matrix is also possible, for example:

```console
vbantxt "Point(ASIO128.IN[2],ASIO128.OUT[1]).dBGain = -8"
```

---

## `Logging`

The --log-level/-l flag allows you to control the verbosity of the application's logging output. 

Acceptable values for this flag are:

- `trace`
- `debug`
- `info`
- `warn`
- `error`
- `fatal`
- `panic`

For example, to set the log level to `debug`, you can use:

```console
vbantxt --log-level=debug "bus[0].eq.on=1 bus[1].gain=-12.8"
```

The default log level is `warn` if the flag is not specified.


## License

`vbantxt` is distributed under the terms of the [MIT](https://spdx.org/licenses/MIT.html) license.
