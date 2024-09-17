![](https://github.com/archway-network/archway/blob/main/banner.png)

# Archway

[![Version](https://img.shields.io/github/v/tag/archway-network/archway.svg?sort=semver&style=flat-square)](https://github.com/archway-network/archway/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/archway-network/archway)
[![Go Report Card](https://goreportcard.com/badge/github.com/archway-network/archway)](https://goreportcard.com/report/github.com/archway-network/archway)
[![codecov](https://codecov.io/gh/archway-network/archway/branch/master/graph/badge.svg)](https://codecov.io/gh/archway-network/archway)
[![License:Apache-2.0](https://img.shields.io/github/license/archway-network/archway.svg?style=flat-square)](https://github.com/archway-network/archway/LICENSE)

The core implementation of the Archway protocol leverages the [Cosmos SDK](https://cosmos.network) and [CosmWasm](https://cosmwasm.com) to reward validators and creators for their contributions to the network.

## System Requirements

The following specifications have been found to work well:

- An x86-64 (amd64) multi-core CPU (AMD / Intel);
  - Higher clock speeds are preferred as CometBFT is mostly single-threaded;
- 64GB RAM;
- 1TB NVMe SSD Storage (disk i/o is crucial);
- 100Mbps bi-directional Internet connection;

## Software Dependencies

The following software should be installed on the target system:

- The Go Programming Language (<https://go.dev>)
- Git Distributed Version Control (<https://git-scm.com>)
- Docker (<https://www.docker.com>)
- GNU Make (<https://www.gnu.org/software/make>)

## Build from Source

[Clone the repository](https://github.com/archway-network/archway), checkout the `main` branch and build:

```sh
cd archway
git checkout main
make install
```

This will install the `archwayd` binary to your `GOPATH`.

## Dockerized Containers

A docker image for production purposes (no shell access):

[Packages: archwayd](https://github.com/orgs/archway-network/packages/container/package/archwayd)

A docker image is also provided for test setups (shell access):

[Packages: archwayd-debug](https://github.com/orgs/archway-network/packages/container/package/archwayd-dev)

## Running localnet

There are two ways to run a localnet, local and containerized

### Containerized

This solution uses docker-compose and docker on the backend.
To setup new localnet use:

```
make localnet
```

To continue last localnet used:

```
make localnet-continue
```

### Local

To set up a localnet natively, you can run the following script:

```sh
./scripts/localnet.sh
```

This script initializes and starts the localnet.

### Accessing Wallets and Balances

When running a localnet, the default mnemonics for the wallets with all the required balance are defined in the `scripts/localnet.sh` file. Here are the mnemonics for the main wallets:

- Validator: 
  ```
  guard cream sadness conduct invite crumble clock pudding hole grit liar hotel maid produce squeeze return argue turtle know drive eight casino maze host
  ```
- Developer:
  ```
  friend excite rough reopen cover wheel spoon convince island path clean monkey play snow number walnut pull lock shoot hurry dream divide concert discover
  ```
- User:
  ```
  any giant turtle pioneer frequent frown harvest ancient episode junior vocal rent shrimp icon idle echo suspect clean cage eternal sample post heavy enough
  ```

These wallets will have sufficient balance to perform testing regarding transactions and deployments of contracts.

## Documentation

To learn more, please [visit the official Archway documentation](https://docs.archway.io).