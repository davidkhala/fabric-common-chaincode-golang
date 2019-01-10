package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (cc CommonChaincode) GetPrivateData(collection, key string) ([]byte) {
	var r, err = cc.CCAPI.GetPrivateData(collection, key)
	PanicError(err)
	return r
}
func (cc CommonChaincode) GetPrivateObj(collection, key string, v interface{}) bool {
	var r, err = cc.CCAPI.GetPrivateData(collection, key)
	PanicError(err)
	if r == nil {
		return false
	}
	FromJson(r, v)
	return true
}
func (cc CommonChaincode) PutPrivateObj(collection, key string, v interface{}) {
	var err = cc.CCAPI.PutPrivateData(collection, key, ToJson(v))
	PanicError(err)
}
func (cc CommonChaincode) PutPrivateData(collection, key string, value []byte) {
	var err = cc.CCAPI.PutPrivateData(collection, key, value)
	PanicError(err)
}

func (cc CommonChaincode) GetPrivateDataByPartialCompositeKey(collection, objectType string, keys []string) (shim.StateQueryIteratorInterface) {
	var r, err = cc.CCAPI.GetPrivateDataByPartialCompositeKey(collection, objectType, keys)
	PanicError(err)
	return r
}
func (cc CommonChaincode) GetPrivateDataByRange(collection, startKey, endKey string) (shim.StateQueryIteratorInterface) {
	var r, err = cc.CCAPI.GetPrivateDataByRange(collection, startKey, endKey)

	PanicError(err)
	return r
}
func (cc CommonChaincode) GetPrivateDataQueryResult(collection, query string) (shim.StateQueryIteratorInterface) {
	var r, err = cc.CCAPI.GetPrivateDataQueryResult(collection, query);
	PanicError(err)
	return r;
}
func (cc CommonChaincode) DelPrivateData(collection, key string) {
	var err = cc.CCAPI.DelPrivateData(collection, key)
	PanicError(err)
}

func (cc CommonChaincode) WorldStatesPrivate(collection, objectType string) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	if objectType == "" {
		keysIterator = cc.GetPrivateDataByRange(collection, "", "")
	} else {
		keysIterator = cc.GetPrivateDataByPartialCompositeKey(collection, objectType, nil)
	}
	return ParseStates(keysIterator)
}
// transaction should be commit to take effect TODO FIXME https://jira.hyperledger.org/browse/FAB-5094
func (cc CommonChaincode)EnableHistoryForPrivateKey(key string) {
	//crypto.HashSha256([]byte(key))
}