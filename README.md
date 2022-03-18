# Charmeleon
A pokedex over SSH project.
![Pokedex in your terminal!](https://raw.githubusercontent.com/SHA65536/Charmeleon/main/startscreen.png)

## Installation
Clone the repository to get the server and pokemon data:
```bash
git clone https://github.com/sha65536/charmeleon
cd charmeleon
go build .
```
## Usage
Charmeleon has two commands:

**Start**: Starts the server on specified address and port.
```
usage: charmeleon start [-h|--help] [-a|--address "<value>"] [-p|--port
                  <integer>]

                  Start the SSH server

Arguments:

  -h  --help     Print help information
  -a  --address  Address to listen on. Default: 0.0.0.0
  -p  --port     Port to listen on. Default: 23234

```
**Build**: Redownloads pokemon sprites, and converts to terminal printable format. **WARNING!!!** This requires your system to have img2xterm installed and imagemagick!
```
usage: charmeleon build [-h|--help]

                  Download images and remake cow files

Arguments:

  -h  --help  Print help information

```

## Made with Charm
[![https://charm.sh](https://stuff.charm.sh/charm/charm-header.png)](https://charm.sh)