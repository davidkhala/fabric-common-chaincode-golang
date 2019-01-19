package golang

import (
	"encoding/json"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
)

type CommonChaincode struct {
	Mock    bool
	Debug   bool
	Name    string
	Logger  *shim.ChaincodeLogger
	Channel string
	CCAPI   shim.ChaincodeStubInterface //chaincode API
}

func (cc *CommonChaincode) Prepare(ccAPI shim.ChaincodeStubInterface) {
	cc.CCAPI = ccAPI
	cc.Channel = ccAPI.GetChannelID()
}
func (cc *CommonChaincode) SetLogger(ccName string) {
	cc.Name = ccName
	cc.Logger = shim.NewLogger(ccName)
}
func (cc CommonChaincode) WorldStates(objectType string, filter func(StateKV) bool) []StateKV {
	var keysIterator shim.StateQueryIteratorInterface
	if objectType == "" {
		keysIterator = cc.GetStateByRange("", "")
	} else {
		keysIterator = cc.GetStateByPartialCompositeKey(objectType, nil)
	}

	return ParseStates(keysIterator, filter)
}

func (cc CommonChaincode) ModifyValue(key string, modifier Modifier, v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		var invalidPtr = json.InvalidUnmarshalError{reflect.TypeOf(v)}
		PanicError(&invalidPtr)
	}
	cc.GetStateObj(key, v)
	modifier(v)
	cc.PutStateObj(key, v)
}

//return empty for if no record.
func (cc CommonChaincode) GetChaincodeID() string {
	var iterator, _ = cc.GetStateByRangeWithPagination("", "", 1, "")
	if !iterator.HasNext() {
		return ""
	}
	var kv, err = iterator.Next()
	PanicError(err)
	return kv.GetNamespace()
}
