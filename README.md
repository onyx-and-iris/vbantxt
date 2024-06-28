![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

# VBAN Sendtext CLI Utility

Send Voicemeeter string requests over a network or to Matrix

## Tested against

- Basic 1.0.8.4
- Banana 2.0.6.4
- Potato 3.0.2.4
- Matrix 1.0.0.3

## Requirements

- [Voicemeeter](https://voicemeeter.com/) or [Matrix](https://vb-audio.com/Matrix/)
- Go 1.18 or greater (if you want to compile yourself, otherwise check `Releases`)

---

## `Command Line`

Pass `host`, `port` and `streamname` as flags, for example:

```
vbantxt-cli -h="gamepc.local" -p=6980 -s=Command1 "strip[0].mute=1 strip[1].mono=1"
```

You may also store them in a `config.toml` located in `home directory / .vbantxt_cli /`

A valid `config.toml` might look like this:

```toml
[connection]
Host="gamepc.local"
Port=6980
Streamname="Command1"
```

- `host` defaults to "localhost"
- `port` defaults to 6980
- `streamname` defaults to "Command1"

Command line flags will override values in a config.toml.

---

## `Script files`

The vbantxt-cli utility accepts a single string request or an array of string requests. This means you can pass scripts stored in files.

For example, in Windows with Powershell you could:

`vbantxt-cli $(Get-Content .\script.txt)`

Or with Bash:

`cat script.txt | xargs vbantxt-cli`

to load commands from a file:

```
strip[0].mute=1;strip[0].mono=0
strip[1].mute=0;strip[1].mono=1
bus[3].eq.On=0
```

---

## `Matrix`

Sending commands to VB-Audio Matrix is also possible, for example:

```
vbantxt-cli -s=streamname "Point(ASIO128.IN[2],ASIO128.OUT[1]).dBGain = -8"
```

---

## `Logging`

Log level may be set by passing the `-l` flag with a number from 0 up to 6 where

0 = Panic, 1 = Fatal, 2 = Error, 3 = Warning, 4 = Info, 5 = Debug, 6 = Trace

Log level defaults to Warning level.
