🔴WARNING: This has been deprecated, please read [this.](https://docs.djtx.network/build/tools/deprecating-ortelius) 🔴

# Ortelius

A data processing pipeline for the [Avalanche network](https://djtx.network).

## Features

- Maintains a persistent log of all consensus events and decisions made on the Avalanche network.
- Indexes Exchange (X), Platform (P), and Contract (C) chain transactions.
- An [API](https://docs.djtx.network/build/tools/ortelius) allowing easy exploration of the index.

## Prerequisite

https://docs.docker.com/engine/install/ubuntu/

https://docs.docker.com/compose/install/

## Quick Start with Standalone Mode on Fuji (testnet) network

The easiest way to get started is to try out the standalone mode.

```shell script
git clone https://github.com/lasthyphen/ortelius.git $GOPATH/github.com/lasthyphen/ortelius
cd $GOPATH/github.com/lasthyphen/ortelius
make dev_env_start
make standalone_run
```

## [Production Deployment](docs/deployment.md)

