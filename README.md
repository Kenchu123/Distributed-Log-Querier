# Distributed Log Querier

UIUC CS425, Distributed Systems: Fall 2023 Machine Programming 1

## Description

See [MP Document](./docs/MP1.CS425.FA23.pdf)

## Installation

### Prerequisites

- Go 1.20

### Build

```bash
# Build the project
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
    id: "1"
  - hostname: "m2"
    port: "7122"
    id: "2"
# ...
```

### Run

```bash
# Run server on each machine
./bin/ds-grep-server [args]
# Usage of ds-grep-server:
#   --port string
#         port to listen on (default "7122")

# Run client on any machine
./bin/ds-grep [args]
# Usage: ds-grep [DS_OPTION]... [OPTION]... PATTERN [FILE]...
# Search for PATTERN in each FILE.

# Example for grep logs/machine.{i}.log
# ./bin/ds-grep --config=.dsgrep/config.yml --machine=.* --machine-ilog --machine-ilog-folder=logs -c GET

# Example for grep logs/machine.log from machine [1-3]
# ./bin/ds-grep --machine=[1-3] -c GET logs/machine.log

# Distributed grep options:
#     --config=PATH  path to config file (default ".dsgrep/config.yml")
#     --machine=REGEX  regex to match machine names
#     --machine-ilog  append machine.$i.log to grep file path
#     --machine-ilog-folder=PATH  folder to store machine.$i.log (default "logs")
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
4. Set `test/[basic,fault,large]/logs/machine.{i}.log` on each VM.

#### Run tests

```bash
# Enter VM
ssh [netid]@[vm-hostname]

# Run tests
$ go test [-v] [-count=1] ./...
```

## Run on boot for VMs

```bash
sudo cp ./bin/ds-grep-server /usr/local/bin
sudo cp ds-grep-server.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable ds-grep-server.service
sudo systemctl start ds-grep-server.service
```

## Contributor

- [Che-Kuang Chu](https://github.com/Kenchu123)
- [Jhih-Wei Lin](https://github.com/williamlin0825)
