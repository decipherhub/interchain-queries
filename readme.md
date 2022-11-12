# Notice
**[This repository is not maintained now]**

Interchain query require interchain query module in appchain and relayer. This repository is implemented for interchain query module, but not adopted.

But **now we still implement [relayer](https://github.com/validance/ibc-rs)** and that repository is not compatible with this repository.

We archived [relayer](https://github.com/decipherhub/ibc-rs) which is compatbile with this repository. So you can still execute this module for intherchainquery.

**[Not satisfy all requirements from spec]**

This implementation doesn't implement validate query proof.



# interchainqueries
**interchainqueries** houses the code for implementing Interchain Query.
The repo is currently a WIP and targeting v1 of cross chain queries.
For more details on the Cross chain queries, take a look at the [specification](https://github.com/cosmos/ibc/tree/main/spec/app/ics-031-crosschain-queries)

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
Unit Test files are located in same directory as the file being tested.
File name convention is ```<file being tested>_test.go```

- x/ibc_query/keeper/grpc_query_test.go
- x/ibc_query/keeper/msg_server_test.go

we use these functionalities for testing
- [Mock external keeper](https://github.com/decipherhub/interchain-queries/blob/main/testutil/keeper/mocks.go)
- [helper function](https://github.com/decipherhub/interchain-queries/blob/main/testutil/keeper/unit_test_helpers.go)
### End to End (e2e) Tests - only appchain
Interchain Query doesn't need IBC connection. So **E2E tests doesn't test simulated connection, packet relays, etc**.\
**It covers only querying chain**. Because Querying to queried chain is implemented in [Relayer](https://github.com/decipherhub/ibc-rs). Just validating status of querying chain in specific scenario.

**Coverage**
- Receive cross chain query request & event emit
- Receive cross chain query result from relayer
- Retrieve and remove result

For E2E test, we use these functionalities
- [setup simapp](https://github.com/decipherhub/interchain-queries/blob/main/testutil/simapp/simapp.go)
- [IBC testing package](https://github.com/cosmos/ibc-go/tree/main/testing)

### Integration Tests - test with relayer implementation
Integration tests run querying, queried chain binaries and relayer binary in docker.\
So it can test real E2E - including relayer functionality(querying to queried chain and return to querying chain)

Integration Tests doesn't be implemented in our repo.  
Link: [e2e test repo](https://github.com/validance/cross-chain-query-e2e/tree/test/add-wallet-account)

### Running Tests

**Unit Test, E2E Test**

```make test```

**Integration Test**
please refer to [e2e test repo](https://github.com/validance/cross-chain-query-e2e/tree/test/add-wallet-account)


## Learn more

- [interchainqueries architecture](./docs/architecture.md)
- [interchainqueries implementation](./docs/implementaion.md)