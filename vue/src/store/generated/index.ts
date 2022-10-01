// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import InterchainQueriesIbcQueryV1 from './interchain_queries.ibc_query.v1'


export default { 
  InterchainQueriesIbcQueryV1: load(InterchainQueriesIbcQueryV1, 'interchain_queries.ibc_query.v1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}