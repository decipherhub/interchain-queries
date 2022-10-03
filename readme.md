# interchainqueries
**interchainqueries** houses the code for implementing Interchain Query.
The repo is currently a WIP and targetting v1 of crosschain queries.
For mor details on the Crosschain queries, take a look at the [specification](https://github.com/cosmos/ibc/tree/main/spec/app/ics-031-crosschain-queries)

## Instructions
**Prerequisites**

```bash
## For OSX or Linux

# go 1.18 (https://formulae.brew.sh/formula/go)
brew install go@1.18
# jq (optional, for testnet) (https://formulae.brew.sh/formula/jq)
brew install jq
# docker (optional, for integration tests, testnet) (https://docs.docker.com/get-docker/)

```

**Installing and running binaries**

```bash
# install interchain-queriesd
make install
# run interchain-queriesd
interchain-queriesd
# (if the above fail, ensure ~/go/bin on $PATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

**Make proto**

```
make proto-all
```

Inspect the [Makefile](./Makefile) if curious.

## Testing

### Unit Tests

preparing..

### End to End (e2e) Tests

preparing..

### Differential Tests (WIP)

preparing..

### Integration Tests

preparing..

### Running Tests

preparing..

## Learn more

- [IBC Docs](https://docs.cosmos.network/master/ibc/)
- [IBC Protocol](https://ibcprotocol.org/)
- [IBC Specs](https://github.com/cosmos/ibc)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)