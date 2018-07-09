package golang

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"errors"
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
	defer iterator.Close();
	var kvs []KVJson
	for iterator.HasNext() {
		kv, err := iterator.Next();
		PanicError(err)
		kvs = append(kvs, KVJson{kv.Namespace, kv.Key, string(kv.Value)})
	}
	return kvs
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
	PanicError(err);
	return key
}

func GetState(ccAPI shim.ChaincodeStubInterface, key string) []byte {
	var bytes, err = ccAPI.GetState(key)
	PanicError(err);
	return bytes
}
func PutState(ccAPI shim.ChaincodeStubInterface, key string, value []byte) {
	var err = ccAPI.PutState(key, value)
	PanicError(err);
}

type KeyModification struct {
	TxId      string
	Value     string
	Timestamp string
	IsDelete  bool
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
	return result;
}
func GetHistoryForKey(ccAPI shim.ChaincodeStubInterface, key string) (shim.HistoryQueryIteratorInterface) {
	var r, err = ccAPI.GetHistoryForKey(key)
	PanicError(err)
	return r;
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
