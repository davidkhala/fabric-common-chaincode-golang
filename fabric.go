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

func WorldStates(ccAPI shim.ChaincodeStubInterface, objectType string) ([]KVJson) {
	var keysIterator shim.StateQueryIteratorInterface
	if objectType == "" {
		keysIterator = GetStateByRange(ccAPI, "", "")
	} else {
		keysIterator = GetStateByPartialCompositeKey(ccAPI, objectType, nil)
	}

	var kvs = ParseStates(keysIterator)
	return kvs
}
func ParseStates(iterator shim.StateQueryIteratorInterface) ([]KVJson) {
	defer iterator.Close()
	var kvs []KVJson
	for iterator.HasNext() {
		kv, err := iterator.Next()
		PanicError(err)
		kvs = append(kvs, KVJson{kv.Namespace, kv.Key, string(kv.Value)})
	}
	return kvs
}

type Modifier func(interface{})

func ModifyValue(ccApi shim.ChaincodeStubInterface, key string, modifier Modifier, v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		PanicString(json.InvalidUnmarshalError{reflect.TypeOf(v)}.Error())
	}
	GetStateObj(ccApi, key, v)
	modifier(v)
	PutStateObj(ccApi, key, v)
}

type KVJson struct {
	Namespace string
	Key       string
	Value     string
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
func GetStateObj(ccAPI shim.ChaincodeStubInterface, key string, v interface{}) {
	var bytes = GetState(ccAPI, key)
	FromJson(bytes, v)
}
func PutStateObj(ccAPI shim.ChaincodeStubInterface, key string, v interface{}) {
	var bytes = ToJson(v)
	PutState(ccAPI, key, bytes)
}
func PutState(ccAPI shim.ChaincodeStubInterface, key string, value []byte) {
	var err = ccAPI.PutState(key, value)
	PanicError(err)
}

type KeyModification struct {
	TxId      string
	Value     string
	Timestamp string
	IsDelete  bool
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
func GetThisMsp(ccApi shim.ChaincodeStubInterface) string {
	var creatorBytes, err = ccApi.GetCreator()
	PanicError(err)
	var creator Creator
	creator, err = ParseCreator(creatorBytes)
	PanicError(err)
	return creator.Msp
}

func ParseHistory(iterator shim.HistoryQueryIteratorInterface) (result []KeyModification) {
	defer iterator.Close()
	for iterator.HasNext() {
		keyModification, err := iterator.Next()
		PanicError(err)
		var timeStamp = keyModification.Timestamp
		var time = timeStamp.Seconds*1000 + int64(timeStamp.Nanos/1000000)
		var translated = KeyModification{
			keyModification.TxId,
			string(keyModification.Value),
			string(time),
			keyModification.IsDelete}
		result = append(result, translated)
	}
	return result
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
			err = errors.New("Unknown panic")
		}
		response.Status = shim.ERROR
		response.Message = err.(error).Error()
	}
}
