import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitPruneCrossChainQueryResult } from "./types/ibc/applications/ibc_query/v1/tx";
import { MsgSubmitCrossChainQuery } from "./types/ibc/applications/ibc_query/v1/tx";
import { MsgSubmitCrossChainQueryResult } from "./types/ibc/applications/ibc_query/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/ibc.applications.ibc_query.v1.MsgSubmitPruneCrossChainQueryResult", MsgSubmitPruneCrossChainQueryResult],
    ["/ibc.applications.ibc_query.v1.MsgSubmitCrossChainQuery", MsgSubmitCrossChainQuery],
    ["/ibc.applications.ibc_query.v1.MsgSubmitCrossChainQueryResult", MsgSubmitCrossChainQueryResult],
    
];

export { msgTypes }