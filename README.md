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

- Run backend server.

    ```sh
    go run ./cmd/bro-backend/main.go
    ```

    You will be prompted to input the listen address if there is no config file.
    The config file will be located in `$HOME/.bro/backend_config.json`.

### CLI Client

- Run CLI client.

    ```sh
    go run ./cmd/bro/main.go --help
    ```

    You will be prompted to input the backend address if there is no config file.
    The config file will be located in `$HOME/.bro/config.json`.
