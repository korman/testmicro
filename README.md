# Testmicro Service

This is the Testmicro service

Generated with

```
micro new github.com/korman/testmicro --namespace=go.micro --type=srv
```

## Getting Started

- [Testmicro Service](#testmicro-service)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.testmicro
- Type: srv
- Alias: testmicro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./testmicro-srv
```

Build a docker image
```
make docker
```