# Fork notice

This is a fork to be used by https://github.com/warden-protocol/wardenprotocol.

All the changes are currently being made to the [release/v1.x.x](https://github.com/warden-protocol/connect/tree/release/v1.x.x) branch, as it's the one we are using at the moment.

---

Original README below:

# Connect [End of Life -- Please fork to use or contribute]

<!-- markdownlint-disable MD013 -->
<!-- markdownlint-disable MD041 -->

[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#wip)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/skip-mev/connect)
[![Go Report Card](https://goreportcard.com/badge/github.com/skip-mev/connect?style=flat-square)](https://goreportcard.com/report/github.com/skip-mev/connect)
[![Version](https://img.shields.io/github/tag/skip-mev/connect.svg?style=flat-square)](https://github.com/skip-mev/connect/releases/latest)
[![Lines Of Code](https://img.shields.io/tokei/lines/github/skip-mev/connect?style=flat-square)](https://github.com/skip-mev/connect)

A general purpose price oracle leveraging ABCI++. Please visit our [docs](https://docs.skip.build/connect/introduction) page for more information!

Connect uses Vote Extensions to create a hyperperformant, extremely secure mechanism for aggregating off-chain data onto a blockchain. It is used by
many of the highest-performance decentralized applications today. If you would like to integrate Connect to power your use case, please contact us on our
[Discord](https://discord.gg/PeBGE9jrbu).

> [!NOTE]
> Connect is **business-licensed software** under BSL, meaning it requires a license to use or reference. It is source viewable, but [**reach out to us on Discord**](https://skip.build/discord) if you are interested in integrating! We are limiting the number of chains we work with to seven in 2024. We apologize if we run out of capacity.

## Install

```shell
$ go install github.com/skip-mev/connect/v2
```

## Overview

The connect repository is composed of the following core packages:

* **abci** - This package contains the [vote extension](./abci/ve/README.md), [proposal](./abci/proposals/README.md), and [preblock handlers](./abci/preblock/oracle/README.md) that are used to broadcast oracle data to the network and to store it in the blockchain.
* **oracle** - This [package](./oracle/) contains the main oracle that aggregates external data sources before broadcasting it to the network. You can reference the provider documentation [here](providers/README.md) to get a high level overview of how the oracle works.
* **providers** - This package contains a collection of [websocket](./providers/websockets/README.md) and [API](./providers/apis/README.md) based data providers that are used by the oracle to collect external data.
* **x/oracle** - This package contains a Cosmos SDK module that allows you to store oracle data on a blockchain.
* **x/marketmap** - This [package](./x/marketmap/README.md) contains  a Cosmos SDK module that allows for market configuration to be stored and updated on a blockchain.

## Validator Usage

To read how to run the oracle as a validator based on the chain, please reference the [validator documentation](https://docs.skip.build/connect/validators/quickstart).

## Developer Usage

To run the oracle, run the following command.

```bash
$ make start-all-dev
```

This will:

1. Start a blockchain with a single validator node. It may take a few minutes to build and reach a point where vote extensions can be submitted.
2. Start the oracle side-car that will aggregate prices from external data providers and broadcast them to the network. To check the current aggregated prices on the side-car, you can run `curl localhost:8080/connect/oracle/v1/prices`.
3. Host a prometheus instance that will scrape metrics from the oracle sidecar. Navigate to http://localhost:9091 to see all network traffic and metrics pertaining to the oracle sidecar. Navigate to http://localhost:8002 to see all application-side oracle metrics.
4. Host a profiler that will allow you to profile the oracle side-car. Navigate to http://localhost:6060 to see the profiler.
5. Host a grafana instance that will allow you to visualize the metrics scraped by prometheus. Navigate to http://localhost:3000 to see the grafana dashboard. The default username and password are `admin` and `admin`, respectively.

After a few minutes, run the following commands to see the prices written to the blockchain:

```bash
# access the blockchain container
$ docker exec -it compose-blockchain-1 bash

# query the price of bitcoin in USD on the node
$ (compose-blockchain-1) ./build/connectd q oracle price BTC USD
```

Result:

```bash
decimals: "8"
id: "0"
nonce: "44"
price:
  block_height: "46"
  block_timestamp: "2024-01-29T01:43:48.735542Z"
  price: "4221100000000"
```

To stop the oracle, run the following command:

```bash
$ make stop-all-dev
```

## Metrics

We have an extensive suite of metrics available to validators and chain operators.
 Please [join our discord](https://discord.gg/PeBGE9jrbu) if you want help setting them up!

Metrics relevant to the oracle service's health + operation are detailed in our docs, [here](https://docs.skip.build/connect/metrics/overview)!

