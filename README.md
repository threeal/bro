# Threeal Bot

[![build status](https://img.shields.io/github/actions/workflow/status/threeal/threeal-bot/build.yml?branch=main)](https://github.com/threeal/threeal-bot/actions/workflows/build.yml)

A personal multi-purpose bot written in [Go](https://go.dev/).

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
