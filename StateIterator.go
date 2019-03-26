package golang

import (
	"fmt"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func parse100States(iterator shim.StateQueryIteratorInterface, filter func(StateKV) bool) ([]StateKV, string) {
	defer PanicError(iterator.Close())
	var kvs []StateKV
	var index = 0
	var lastKey = ""
	for iterator.HasNext() {
		fmt.Println("[debug]index", index)
		if index >= 100 {
			return kvs, lastKey
		}

		kv, err := iterator.Next()
		PanicError(err)
		lastKey = kv.Key
		var stateKV = StateKV{kv.Namespace, kv.Key, string(kv.Value)}
		if filter == nil || filter(stateKV) {
			kvs = append(kvs, stateKV)
		}
		index++

	}
	return kvs, ""
}
func (cc CommonChaincode) WorldStatesPrivate(collection, startKey string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetPrivateDataByRange(collection, startKey, "")

	var states, bookmark = parse100States(keysIterator, filter)
	if bookmark != "" {
		return append(states, cc.WorldStatesPrivate(collection, bookmark, filter)...)
	} else {
		return states
	}
}

func (cc CommonChaincode) WorldStates(startKey string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	keysIterator = cc.GetStateByRange(startKey, "")
	var states, bookmark = parse100States(keysIterator, filter)
	if bookmark != "" {
		return append(states, cc.WorldStates(bookmark, filter)...)
	} else {
		return states
	}
}
