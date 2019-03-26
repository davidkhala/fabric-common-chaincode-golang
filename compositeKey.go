package golang

import "github.com/hyperledger/fabric/core/chaincode/shim"

//TODO 100 iterator limit
func (cc CommonChaincode) WorldStatesPrivateComposite(collection, objectType string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetPrivateDataByPartialCompositeKey(collection, objectType, nil)
	return ParseStates(keysIterator, filter)
}

//TODO 100 iterator limit
func (cc CommonChaincode) WorldStatesComposite(objectType string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetStateByPartialCompositeKey(objectType, nil)
	return ParseStates(keysIterator, filter)
}
