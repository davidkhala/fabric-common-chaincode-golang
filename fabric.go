package golang

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"errors"
	"time"
	"fmt"
	"reflect"
	"encoding/json"
)

func WorldStates(ccAPI shim.ChaincodeStubInterface, objectType string) (States) {
	var keysIterator shim.StateQueryIteratorInterface
	if objectType == "" {
		keysIterator = GetStateByRange(ccAPI, "", "")
	} else {
		keysIterator = GetStateByPartialCompositeKey(ccAPI, objectType, nil)
	}

	var state States
	state.ParseStates(keysIterator)
	return state
}

func ModifyValue(ccApi shim.ChaincodeStubInterface, key string, modifier Modifier, v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		var invalidPtr = json.InvalidUnmarshalError{reflect.TypeOf(v)}
		PanicString(invalidPtr.Error())
	}
	GetStateObj(ccApi, key, v)
	modifier(v)
	PutStateObj(ccApi, key, v)
}

func SplitCompositeKey(ccAPI shim.ChaincodeStubInterface, compositeKey string) (string, []string) {
	objectType, attributes, err := ccAPI.SplitCompositeKey(compositeKey)
	PanicError(err)
	return objectType, attributes
}
func CreateCompositeKey(ccAPI shim.ChaincodeStubInterface, objectType string, attributes []string) string {
	var key, err = ccAPI.CreateCompositeKey(objectType, attributes)
	PanicError(err)
	return key
}

func GetState(ccAPI shim.ChaincodeStubInterface, key string) []byte {
	var bytes, err = ccAPI.GetState(key)
	PanicError(err)
	return bytes
}
func GetStateObj(ccAPI shim.ChaincodeStubInterface, key string, v interface{}) bool {
	var bytes = GetState(ccAPI, key)
	if bytes == nil {
		return false
	}
	FromJson(bytes, v)
	return true
}
func PutStateObj(ccAPI shim.ChaincodeStubInterface, key string, v interface{}) {
	var bytes = ToJson(v)
	PutState(ccAPI, key, bytes)
}
func PutState(ccAPI shim.ChaincodeStubInterface, key string, value []byte) {
	var err = ccAPI.PutState(key, value)
	PanicError(err)
}

func GetTxTime(ccApi shim.ChaincodeStubInterface) (time.Time) {
	ts, err := ccApi.GetTxTimestamp()
	PanicError(err)

	var t time.Time
	if ts == nil {
		t = time.Unix(0, 0).UTC() // treat nil like the empty Timestamp
	} else {
		t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
	}

	if ts.Seconds < minValidSeconds {
		PanicError(fmt.Errorf("timestamp: %v before 0001-01-01", ts))
	}
	if ts.Seconds >= maxValidSeconds {
		PanicError(fmt.Errorf("timestamp: %v after 10000-01-01", ts))
	}
	if ts.Nanos < 0 || ts.Nanos >= 1e9 {
		PanicError(fmt.Errorf("timestamp: %v: nanos not in range [0, 1e9)", ts))
	}
	return t

}
func GetThisCreator(ccApi shim.ChaincodeStubInterface) Creator {
	var creatorBytes, err = ccApi.GetCreator()
	PanicError(err)
	var creator Creator
	creator, err = ParseCreator(creatorBytes)
	PanicError(err)
	return creator
}

func GetHistoryForKey(ccAPI shim.ChaincodeStubInterface, key string) (shim.HistoryQueryIteratorInterface) {
	var r, err = ccAPI.GetHistoryForKey(key)
	PanicError(err)
	return r
}
func GetStateByPartialCompositeKey(ccAPI shim.ChaincodeStubInterface, objectType string, keys []string) shim.StateQueryIteratorInterface {
	var r, err = ccAPI.GetStateByPartialCompositeKey(objectType, keys)
	PanicError(err)
	return r
}
func GetStateByRange(ccAPI shim.ChaincodeStubInterface, startKey string, endKey string) shim.StateQueryIteratorInterface {
	var r, err = ccAPI.GetStateByRange(startKey, endKey)
	PanicError(err)
	return r
}
func PanicDefer(response *peer.Response) {
	if err := recover(); err != nil {
		switch x := err.(type) {
		case string:
			err = errors.New(x)
		case error:
		default:
			err = errors.New("unknown panic")
		}
		fmt.Println(err)
		response.Status = shim.ERROR
		response.Message = err.(error).Error()
	}
}

type CommonChaincode struct {
	Mock  bool
	Debug bool
	CCAPI *shim.ChaincodeStubInterface //chaincode API
}

func (cc *CommonChaincode) Prepare(ccAPI *shim.ChaincodeStubInterface) {
	cc.CCAPI = ccAPI
}
