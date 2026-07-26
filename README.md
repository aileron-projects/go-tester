# go-tester

**Go library for testing.**

<div align="center">

[![GoDoc](https://godoc.org/github.com/aileron-projects/go-tester?status.svg)](http://godoc.org/github.com/aileron-projects/go-tester)
[![Test](https://github.com/aileron-projects/go-tester/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/aileron-projects/go-tester/actions/workflows/test.yaml?query=branch%3Amain)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellow.svg)](./LICENSE)

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-tester)
[![OpenSourceInsight](https://badgen.net/badge/open%2Fsource%2F/insight/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-tester)
[![OSS Insight](https://badgen.net/badge/OSS/Insight/orange)](https://ossinsight.io/analyze/aileron-projects/go-tester)

</div>

## Features

- Assertion
- Global variable replacer (os.Stdout, os.Stderr, rand.Reader, etc)
- io.Reader and io.Writer that can return errors

## Tested Environments

Operating System:

- `Linux`: [ubuntu-latest](https://github.com/actions/runner-images)
- `Windows`: [windows-latest](https://github.com/actions/runner-images)
- `macOS`: [macos-latest](https://github.com/actions/runner-images)

Architecture (Using QEMU on linux):

- x86: `amd64`, `386`
- arm: `arm/v5`, `arm/v6`, `arm/v7`, `arm64`
- risc: `riscv64`, `loong64`
- ppc: `ppc64`, `ppc64le`
- mips: `mips`, `mips64`, `mips64le`, `mipsle`
- ibm: `s390x`

## Release Cycle

- Releases are made as needed.
- [Semantic Versioning](https://semver.org/) `vX.Y.Z` is used.

## License

[Apache-2.0](LICENSE)

## Usage

TODO
