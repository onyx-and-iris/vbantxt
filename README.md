![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

# VBAN Sendtext

Send Voicemeeter/Matrix vban requests.

For an outline of past/future changes refer to: [CHANGELOG](CHANGELOG.md)

## Tested against

-   Basic 1.1.1.8
-   Banana 2.1.1.8
-   Potato 3.1.1.8
-   Matrix 1.0.1.2

## Requirements

-   [Voicemeeter](https://voicemeeter.com/) or [Matrix](https://vb-audio.com/Matrix/)
-   Go 1.18 or greater (if you want to compile yourself, otherwise check `Releases`)

---

## `Use`

`go get github.com/onyx-and-iris/vbantxt`

```go
package main

import (
	"log"

	"github.com/onyx-and-iris/vbantxt"
)

func main() {
	var (
		host       string = "vm.local"
		port       int    = 6980
		streamname string = "onyx"
	)

	client, err := vbantxt.New(host, port, streamname)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	err = client.Send("strip[0].mute=0")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
```

## `Command Line`

### Configuration

#### Flags

-	--host/-H: defaults to localhost
-	--port/-p: defaults to 6980
-	--streamname/-s: defaults to Command1
-	--config/-C: defaults to `$XDG_CONFIG_HOME / vbantxt / config.toml`
-	--loglevel/-l: defaults to warn
-	--version/-v: print the vbantxt version and exit

For example:

```console
vbantxt --host="gamepc.local" --port=6980 --streamname=Command1 "strip[0].mute=1 strip[1].mono=1"
```

#### Environment Variables

Load the following values from your environment:

```bash
#!/usr/bin/env bash

export VBANTXT_HOST=localhost
export VBANTXT_PORT=6980
export VBANTXT_STREAMNAME=onyx
```

Flags will override environment variables.

#### TOML Config

By default the config loader will look for a config in:

-	$XDG_CONFIG_HOME / vbantxt / config.toml (see [os.UserConfigDir](https://pkg.go.dev/os#UserConfigDir))

A valid config.toml might look like this:

```toml
host="gamepc.local"
port=6980
streamname="Command1"
```

A custom config path may be passed with the --config/-C flag.

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

The --loglevel/-l flag allows you to control the verbosity of the application's logging output. 

Acceptable values for this flag are:

- `debug`
- `info`
- `warn`
- `error`
- `fatal`
- `panic`

For example, to set the log level to `debug`, you can use:

```console
vbantxt --loglevel=debug "bus[0].eq.on=1 bus[1].gain=-12.8"
```

The default log level is `warn` if the flag is not specified.
