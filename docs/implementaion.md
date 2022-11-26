## Spec
[ics-031-crosschain-queries](https://github.com/cosmos/ibc/tree/main/spec/app/ics-031-crosschain-queries)

## Terminology
- Querying chain: chain which requests cross chain query to relayer
- Queried chain: target chain which querying chain wants to query.

# tx
***MsgSubmitCrossChainQuery***
```Go
type MsgSubmitCrossChainQuery struct {
	Path               string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LocalTimeoutHeight types.Height `protobuf:"bytes,2,opt,name=local_timeout_height,json=localTimeoutHeight,proto3" json:"local_timeout_height" yaml:"timeout_height"`
	LocalTimeoutStamp  uint64       `protobuf:"varint,3,opt,name=local_timeout_stamp,json=localTimeoutStamp,proto3" json:"local_timeout_stamp,omitempty"`
	QueryHeight        uint64       `protobuf:"varint,4,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
	ChainId            string       `protobuf:"bytes,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Sender             string       `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
}
```
request cross chain query
- ```Path```: identify queried module.
- ```LocalTimeoutHeight```: be used to check timeout when cross chain query result arrived. If current block height is bigger, result status becomes *QUERY_RESULT_TIMEOUT*. If set LocalTimeoutHeight 0, ignore checking timeout height.
- ```LocalTimeoutStamp```: be used to check timeout when cross chain query result arrived. If current block timestamp is bigger, result status becomes *QUERY_RESULT_TIMEOUT*. If set LocalTimeoutStamp 0, ignore checking timeout stamp.
- ```QueryHeight```: the height at which the relayer must query the queried chain.
- ```ChainId```: identify the querying chain's client of the queried chain.
- ```Sender```: address of tx sender.

***MsgSubmitCrossChainQueryResponse***
```Go
type MsgSubmitCrossChainQueryResponse struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}
```
Response of ```MsgSubmitCrossChainQuery```.\
Return cross chain query id. User can retrieve cross chain query result by query id.

- ```Id```: cross chain query id.

***MsgSubmitCrossChainQueryResult***
```Go
type MsgSubmitCrossChainQueryResult struct {
	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	QueryHeight uint64           `protobuf:"varint,2,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
	Result      QueryResult      `protobuf:"varint,3,opt,name=result,proto3,enum=ibc.applications.ibc_query.v1.QueryResult" json:"result,omitempty"`
	Data        []byte           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Sender      string           `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	ProofSpecs  []*_go.ProofSpec `protobuf:"bytes,6,rep,name=proof_specs,json=proofSpecs,proto3" json:"proof_specs,omitempty"`
}
```
```MsgSubmitCrossChainQueryResult``` is used by relayer.
Relayer submit cross chain query result to querying chain. 

- ```Id```: identify queried module.
- ```QueryHeight```: the height at which the relayer must query the queried chain.
- ```Result```: status which shows query success, fail, timeout. 
- ```Data```: query result data.
- ```Sender```: address of cross chain query requester(user).
- ```ProofSpecs```: use for checking valid result. ***But proof for query doesn't be implemented in tendermint now.*** 

***MsgSubmitPruneCrossChainQueryResult***
```Go
type MsgSubmitPruneCrossChainQueryResult struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}
```
Retrieve cross chain query result and remove it.\
Check whether ```MsgSubmitPruneCrossChainQueryResult``` sender is same as cross chian query requester.

- ```Id```: cross chain query id.
- ```Sender```: ```MsgSubmitPruneCrossChainQueryResult``` sender.

# query
***QueryCrossChainQuery***
```Go
type QueryCrossChainQuery struct {
	// query id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}
```

Retrieve ```CrossChainQuery``` by query id.\
```CrossChainQuery``` is **cross chain query request**, not result.
- ```Id```: query id

***QueryCrossChainQueryResult***
```Go
type QueryCrossChainQueryResult struct {
	// query id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}
```

Retrieve ```CrossChainQueryResult``` by query id.
- ```Id```: query id


# What's different from spec
## Capability key
Use capability for authenticating authentication to delete cross chain query (```MsgSubmitPruneCrossChainQueryResult```)

In spec, user receive capability of cross chain query as response of sending cross chain query request.
But we implement capaibility logic little different.
 
### make capability key with query id and user id 

spec requires making capability with query id. But it makes all people who know query id can delete query result before query requester saw results.\
So we implement that making capability require both query id and user id.


### Response doesn't return capability key and request doesn't require capabilitiy 

In spec, ```CrossChainQueryRequest``` return capability key and ```PruneCrossChainQueryResult``` require capability key. 
But in our implementation ```MsgSubmitPruneCrossChainQueryResult``` doesn't require capability key. 

## Proof
### We don't implement crosschain query validation
[Relayer implementation](https://github.com/decipherhub/ibc-rs)

In cosmos sdk, there are [three types of query](https://docs.cosmos.network/main/core/grpc_rest). REST, GRPC, abci-query(tendermint).
- REST: Write target query module and target in URI. So relayer just make URI with target info and send to queried chain. **But query doesn't retrieve proof from chain.(Not implemented yet)**
- GRPC: **Relayer should have grpc proto file for target module**. So request to custom module can't be processed if relayer doesn't have proto file.
- abci-query: It can be used like REST but more complicated(require add some prefix). But abci-query is treated as legacy so not recommended to use. 

We implemnt interchain query with [REST](https://docs.cosmos.network/main/core/grpc_rest#rest-server) in relayer. Because using GRPC instead of REST makes module flexibility limited.

But as stated above, REST don't retrieve proof from queried chain. So we don't implement crosschain query validation