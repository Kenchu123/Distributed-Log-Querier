# CS425 Distributed Systems MP1

## Description

See [MP Document](./docs/MP1.CS425.FA23.pdf)

## Installation

### Prerequisites

- Go 1.20

### Build

```bash
# Build project
git clone https://gitlab.engr.illinois.edu/ckchu2/cs425-mp1.git
cd cs425-mp1
go build -o bin/ds-grep cmd/ds-grep/main.go
go build -o bin/ds-grep-server cmd/ds-grep-server/main.go
```

## Usage

### Configuration

Set all machines' hostnames and ports in `.config.yml` file.

```yaml
machines:
  - hostname: "m1"
    port: "7122"
  - hostname: "m2"
    port: "7122"
# ...
```

### Run

```bash
# Run server on each machine
./bin/ds-grep-server [args]
# Usage of ds-grep-server:
#   -config string
#         path to config file (default ".dsgrep/config.yml")

# Run client on any machine
./bin/ds-grep [args]
# Usage: ds-grep [DS_OPTION]... [OPTION]... PATTERN [FILE]...
# Search for PATTERN in each FILE.
# Example: ds-grep -i 'hello world' menu.h main.c

# Distributed grep options:
#     --config=PATH  path to config file (default ".dsgrep/config.yml")
#     --machine=REGEX  regex to match machine names
#
# Options:
#   ... (same as grep)
```

## Testing

### Docker

#### Set test logs

#### Build and Run containers

```bash
docker compose up -d [--build]
```

#### Run tests

```bash
# Enter dev container
docker exec -it dev /bin/ash

# Run tests
$ go test [-v] [-count=1] ./...
```

#### Useful commands

```bash
# See logs
docker compose logs -f [m[1-10]]

# Close
docker compose down
```

### VMs

#### Prerequisites

1. Clone this repo to each VM.
2. Build this project on each VM.
3. Start all `ds-grep-server` on each VM.
4. Set `logs/machine.i.log` on each VM.

#### Run tests

```bash
# Enter VM
ssh [netid]@[vm-hostname]

# Run tests
$ go test [-v] [-count=1] ./...
```

## Contributor

- [Che-Kuang Chu (ckchu2)](https://gitlab.engr.illinois.edu/ckchu2)
- [Jhih-Wei Lin (jhihwei2)](https://gitlab.engr.illinois.edu/jhihwei2)
