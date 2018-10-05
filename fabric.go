package golang

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"reflect"
	"time"
)

const (
	// Seconds field of the earliest valid Timestamp.
	// This is time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	minValidSeconds = -62135596800
	// Seconds field just after the latest valid Timestamp.
	// This is time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	maxValidSeconds = 253402300800
)

func (cc CommonChaincode) WorldStates(objectType string) States {
	var keysIterator shim.StateQueryIteratorInterface
	if objectType == "" {
		keysIterator = cc.GetStateByRange("", "")
	} else {
		keysIterator = cc.GetStateByPartialCompositeKey(objectType, nil)
	}

	var state States
	state.ParseStates(keysIterator)
	return state
}
func (cc CommonChaincode) InvokeChaincode(chaincodeName string, args [][]byte, channel string) peer.Response {
	var resp = cc.CCAPI.InvokeChaincode(chaincodeName, args, channel)
	if resp.Status >= shim.ERRORTHRESHOLD {
		panic(errors.New(string(ToJson(resp))))
	}
	return resp
}

func (cc CommonChaincode) ModifyValue(key string, modifier Modifier, v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		var invalidPtr = json.InvalidUnmarshalError{reflect.TypeOf(v)}
		PanicString(invalidPtr.Error())
	}
	cc.GetStateObj(key, v)
	modifier(v)
	cc.PutStateObj(key, v)
}

func (cc CommonChaincode) SplitCompositeKey(compositeKey string) (string, []string) {
	objectType, attributes, err := cc.CCAPI.SplitCompositeKey(compositeKey)
	PanicError(err)
	return objectType, attributes
}
func (cc CommonChaincode) CreateCompositeKey(objectType string, attributes []string) string {
	var key, err = cc.CCAPI.CreateCompositeKey(objectType, attributes)
	PanicError(err)
	return key
}

func (cc CommonChaincode) GetState(key string) []byte {
	var bytes, err = cc.CCAPI.GetState(key)
	PanicError(err)
	return bytes
}
func (cc CommonChaincode) GetStateObj(key string, v interface{}) bool {
	var bytes = cc.GetState(key)
	if bytes == nil {
		return false
	}
	FromJson(bytes, v)
	return true
}
func (cc CommonChaincode) GetTransient() map[string][]byte {
	transient, err := cc.CCAPI.GetTransient()
	PanicError(err)
	return transient
}
func (cc CommonChaincode) PutStateObj(key string, v interface{}) {
	var bytes = ToJson(v)
	cc.PutState(key, bytes)
}
func (cc CommonChaincode) PutState(key string, value []byte) {
	var err = cc.CCAPI.PutState(key, value)
	PanicError(err)
}

func (cc CommonChaincode) GetTxTime() time.Time {
	ts, err := cc.CCAPI.GetTxTimestamp()
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
func (cc CommonChaincode) GetThisCreator() Creator {
	var creatorBytes, err = cc.CCAPI.GetCreator()
	PanicError(err)
	var creator Creator
	creator, err = ParseCreator(creatorBytes)
	PanicError(err)
	return creator
}

func (cc CommonChaincode) GetHistoryForKey(key string) shim.HistoryQueryIteratorInterface {
	var r, err = cc.CCAPI.GetHistoryForKey(key)
	PanicError(err)
	return r
}
func (cc CommonChaincode) GetStateByPartialCompositeKey(objectType string, keys []string) shim.StateQueryIteratorInterface {
	var r, err = cc.CCAPI.GetStateByPartialCompositeKey(objectType, keys)
	PanicError(err)
	return r
}
func (cc CommonChaincode) GetStateByRange(startKey string, endKey string) shim.StateQueryIteratorInterface {
	var r, err = cc.CCAPI.GetStateByRange(startKey, endKey)
	PanicError(err)
	return r
}
func (cc CommonChaincode) SetEvent(name string, payload []byte) {
	var err = cc.CCAPI.SetEvent(name, payload)
	PanicError(err)
}

func DeferPeerResponse(response *peer.Response) {
	var handler = func(errString string) bool {
		response.Status = shim.ERROR
		response.Message = errString
		return true
	}
	Deferred(handler)
}

type CommonChaincode struct {
	Mock    bool
	Debug   bool
	Name    string
	Logger  *shim.ChaincodeLogger
	Channel string
	CCAPI   shim.ChaincodeStubInterface //chaincode API
	shim.Chaincode
}

func (cc *CommonChaincode) Prepare(ccAPI shim.ChaincodeStubInterface) {
	cc.CCAPI = ccAPI
	cc.Channel = ccAPI.GetChannelID()
}
func (cc *CommonChaincode) SetLogger(ccName string) {
	cc.Name = ccName
	cc.Logger = shim.NewLogger(ccName)
}
func ChaincodeStart(cc shim.Chaincode) {
	var err = shim.Start(cc)
	PanicError(err)
}
