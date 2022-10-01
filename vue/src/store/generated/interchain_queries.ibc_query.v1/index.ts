import { Client, registry, MissingWalletError } from 'cosmos-interchain-queries-client-ts'

import { CrossChainQuery } from "cosmos-interchain-queries-client-ts/interchain_queries.ibc_query.v1/types"
import { CrossChainQueryResult } from "cosmos-interchain-queries-client-ts/interchain_queries.ibc_query.v1/types"
import { IBCQueryPacketData } from "cosmos-interchain-queries-client-ts/interchain_queries.ibc_query.v1/types"


export { CrossChainQuery, CrossChainQueryResult, IBCQueryPacketData };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				CrossChainQueryResult: {},
				
				_Structure: {
						CrossChainQuery: getStructure(CrossChainQuery.fromPartial({})),
						CrossChainQueryResult: getStructure(CrossChainQueryResult.fromPartial({})),
						IBCQueryPacketData: getStructure(IBCQueryPacketData.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getCrossChainQueryResult: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CrossChainQueryResult[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: interchain_queries.ibc_query.v1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryCrossChainQueryResult({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.InterchainQueriesIbcQueryV1.query.queryCrossChainQueryResult( key.id)).data
				
					
				commit('QUERY', { query: 'CrossChainQueryResult', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCrossChainQueryResult', payload: { options: { all }, params: {...key},query }})
				return getters['getCrossChainQueryResult']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCrossChainQueryResult API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSubmitCrossChainQuery({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.InterchainQueriesIbcQueryV1.tx.sendMsgSubmitCrossChainQuery({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitCrossChainQuery:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitCrossChainQuery:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitCrossChainQueryResult({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.InterchainQueriesIbcQueryV1.tx.sendMsgSubmitCrossChainQueryResult({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitCrossChainQueryResult:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitCrossChainQueryResult:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSubmitCrossChainQuery({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.InterchainQueriesIbcQueryV1.tx.msgSubmitCrossChainQuery({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitCrossChainQuery:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitCrossChainQuery:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitCrossChainQueryResult({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.InterchainQueriesIbcQueryV1.tx.msgSubmitCrossChainQueryResult({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitCrossChainQueryResult:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitCrossChainQueryResult:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
