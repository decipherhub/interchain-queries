import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitCrossChainQuery } from "./types/ibc-query/v1/tx";
import { MsgSubmitCrossChainQueryResult } from "./types/ibc-query/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/interchain_queries.ibc_query.v1.MsgSubmitCrossChainQuery", MsgSubmitCrossChainQuery],
    ["/interchain_queries.ibc_query.v1.MsgSubmitCrossChainQueryResult", MsgSubmitCrossChainQueryResult],
    
];

export { msgTypes }