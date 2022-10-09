// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgSubmitPruneCrossChainQueryResult } from "./types/ibc/applications/ibc_query/v1/tx";
import { MsgSubmitCrossChainQuery } from "./types/ibc/applications/ibc_query/v1/tx";
import { MsgSubmitCrossChainQueryResult } from "./types/ibc/applications/ibc_query/v1/tx";


export { MsgSubmitPruneCrossChainQueryResult, MsgSubmitCrossChainQuery, MsgSubmitCrossChainQueryResult };

type sendMsgSubmitPruneCrossChainQueryResultParams = {
  value: MsgSubmitPruneCrossChainQueryResult,
  fee?: StdFee,
  memo?: string
};

type sendMsgSubmitCrossChainQueryParams = {
  value: MsgSubmitCrossChainQuery,
  fee?: StdFee,
  memo?: string
};

type sendMsgSubmitCrossChainQueryResultParams = {
  value: MsgSubmitCrossChainQueryResult,
  fee?: StdFee,
  memo?: string
};


type msgSubmitPruneCrossChainQueryResultParams = {
  value: MsgSubmitPruneCrossChainQueryResult,
};

type msgSubmitCrossChainQueryParams = {
  value: MsgSubmitCrossChainQuery,
};

type msgSubmitCrossChainQueryResultParams = {
  value: MsgSubmitCrossChainQueryResult,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgSubmitPruneCrossChainQueryResult({ value, fee, memo }: sendMsgSubmitPruneCrossChainQueryResultParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitPruneCrossChainQueryResult: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitPruneCrossChainQueryResult({ value: MsgSubmitPruneCrossChainQueryResult.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitPruneCrossChainQueryResult: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSubmitCrossChainQuery({ value, fee, memo }: sendMsgSubmitCrossChainQueryParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitCrossChainQuery: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitCrossChainQuery({ value: MsgSubmitCrossChainQuery.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitCrossChainQuery: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSubmitCrossChainQueryResult({ value, fee, memo }: sendMsgSubmitCrossChainQueryResultParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitCrossChainQueryResult: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitCrossChainQueryResult({ value: MsgSubmitCrossChainQueryResult.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitCrossChainQueryResult: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgSubmitPruneCrossChainQueryResult({ value }: msgSubmitPruneCrossChainQueryResultParams): EncodeObject {
			try {
				return { typeUrl: "/ibc.applications.ibc_query.v1.MsgSubmitPruneCrossChainQueryResult", value: MsgSubmitPruneCrossChainQueryResult.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitPruneCrossChainQueryResult: Could not create message: ' + e.message)
			}
		},
		
		msgSubmitCrossChainQuery({ value }: msgSubmitCrossChainQueryParams): EncodeObject {
			try {
				return { typeUrl: "/ibc.applications.ibc_query.v1.MsgSubmitCrossChainQuery", value: MsgSubmitCrossChainQuery.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitCrossChainQuery: Could not create message: ' + e.message)
			}
		},
		
		msgSubmitCrossChainQueryResult({ value }: msgSubmitCrossChainQueryResultParams): EncodeObject {
			try {
				return { typeUrl: "/ibc.applications.ibc_query.v1.MsgSubmitCrossChainQueryResult", value: MsgSubmitCrossChainQueryResult.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitCrossChainQueryResult: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			IbcApplicationsIbcQueryV1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;