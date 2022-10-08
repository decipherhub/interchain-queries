import { txClient, queryClient, MissingWalletError , registry} from './module'

import { CrossChainQuery } from "./module/types/ibc_query/v1/crosschainquery"
import { CrossChainQueryResult } from "./module/types/ibc_query/v1/crosschainquery"
import { IBCQueryPacketData } from "./module/types/ibc_query/v1/packet"


export { CrossChainQuery, CrossChainQueryResult, IBCQueryPacketData };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				CrossChainQuery: {},
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
				getCrossChainQuery: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CrossChainQuery[JSON.stringify(params)] ?? {}
		},
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
			console.log('Vuex module: ibc_query.v1 initialized!')
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
		
		
		
		 		
		
		
		async QueryCrossChainQuery({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCrossChainQuery( key.id)).data
				
					
				commit('QUERY', { query: 'CrossChainQuery', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCrossChainQuery', payload: { options: { all }, params: {...key},query }})
				return getters['getCrossChainQuery']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCrossChainQuery API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCrossChainQueryResult({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCrossChainQueryResult( key.id)).data
				
					
				commit('QUERY', { query: 'CrossChainQueryResult', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCrossChainQueryResult', payload: { options: { all }, params: {...key},query }})
				return getters['getCrossChainQueryResult']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCrossChainQueryResult API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSubmitCrossChainQuery({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitCrossChainQuery(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
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
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitCrossChainQueryResult(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
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
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitCrossChainQuery(value)
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
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitCrossChainQueryResult(value)
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
