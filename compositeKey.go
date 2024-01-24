package golang

import (
	"github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc CommonChaincode) WorldStatesPrivateComposite(collection, objectType string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetPrivateDataByPartialCompositeKey(collection, objectType, nil)
	return ParseStates(keysIterator, filter)
}

func (cc CommonChaincode) WorldStatesComposite(objectType string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetStateByPartialCompositeKey(objectType, nil)
	return ParseStates(keysIterator, filter)
}
func (cc CommonChaincode) CreateCompositeKey(objectType string, attributes []string) string {
	var key, err = cc.CCAPI.CreateCompositeKey(objectType, attributes)
	goutils.PanicError(err)
	return key
}
func (cc CommonChaincode) GetStateByPartialCompositeKey(objectType string, keys []string) shim.StateQueryIteratorInterface {
	var r, err = cc.CCAPI.GetStateByPartialCompositeKey(objectType, keys)
	goutils.PanicError(err)
	return r
}

// GetStateByPartialCompositeKeyWithPagination This call is only supported in a read only transaction.
func (cc CommonChaincode) GetStateByPartialCompositeKeyWithPagination(objectType string, keys []string, pageSize int, bookmark string) (shim.StateQueryIteratorInterface, QueryResponseMetadata) {
	var iterator, r, err = cc.CCAPI.GetStateByPartialCompositeKeyWithPagination(objectType, keys, int32(pageSize), bookmark)
	goutils.PanicError(err)
	return iterator, QueryResponseMetadata{int(r.FetchedRecordsCount), r.Bookmark}
}

func (cc CommonChaincode) SplitCompositeKey(compositeKey string) (string, []string) {
	objectType, attributes, err := cc.CCAPI.SplitCompositeKey(compositeKey)
	goutils.PanicError(err)
	return objectType, attributes
}
