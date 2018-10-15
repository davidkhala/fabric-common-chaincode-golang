package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// It is only supported for state databases that support rich query, e.g., CouchDB.
func (cc CommonChaincode) GetQueryResult(query string) (shim.StateQueryIteratorInterface) {
	var result, err = cc.CCAPI.GetQueryResult(query)
	PanicError(err)
	return result
}

// It is only supported for state databases that support rich query, e.g., CouchDB.
func (cc CommonChaincode) GetQueryResultWithPagination(query string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata) {
	var result, metadata, err = cc.CCAPI.GetQueryResultWithPagination(query, pageSize, bookmark)
	PanicError(err)
	return result, metadata
}
