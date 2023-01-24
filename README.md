# Bro

[![build status](https://img.shields.io/github/actions/workflow/status/threeal/bro/build.yml?branch=main)](https://github.com/threeal/bro/actions/workflows/build.yml)
[![tests status](https://img.shields.io/testspace/pass-ratio/threeal/threeal:bro/main)](https://threeal.testspace.com/projects/threeal:bro)
[![coverage status](https://img.shields.io/coveralls/github/threeal/bro/main)](https://coveralls.io/github/threeal/bro)

Your friendly, personal, multi-purpose [buddy](https://en.wiktionary.org/wiki/buddy) written in [Go](https://go.dev/).

## Build

- Requirements: [Go](https://go.dev/doc/install), [Protocol Compiler](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) (Protobuf's protoc).
- Generate sources from `*.proto` files.
  ```sh
  go generate ./pkg/...
  ```
- (Optional) Build targets.
  ```sh
  go build ./cmd/...
  ```

## Usage

### Backend Server

- (Optional) Set listen address and port of the TCP server (default is `:50051`).
  ```sh
  export THREEAL_BOT_ADDR=':50052'
  ```
- Run backend server.
  ```sh
  go run ./cmd/backend/main.go
  ```

### CLI Client

- (Optional) Set address and port of the target backend server (default is `localhost:50051`).
  ```sh
  export THREEAL_BOT_ADDR='192.168.10.10:50052'
  ```
  > You must specify this if backend server is not ran locally or not using the default port.
- Run CLI client.
  ```sh
  go run ./cmd/client/main.go --help
  ```
