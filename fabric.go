package golang

import (
	"errors"
	"fmt"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"runtime/debug"
	"time"
)

func (cc CommonChaincode) InvokeChaincode(chaincodeName string, args [][]byte, channel string) peer.Response {
	var resp = cc.CCAPI.InvokeChaincode(chaincodeName, args, channel)
	if resp.Status >= shim.ERRORTHRESHOLD {
		var errorPB = PeerResponse{
			resp.Status,
			resp.Message,
			string(resp.Payload),
		}
		panic(errors.New(string(ToJson(errorPB))))
	}
	return resp
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
func (cc CommonChaincode) GetBinding() []byte {
	var result, err = cc.CCAPI.GetBinding()
	PanicError(err)
	return result
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
func (cc CommonChaincode) DelState(key string) {
	var err = cc.CCAPI.DelState(key)
	PanicError(err)
}
func (cc CommonChaincode) GetTxTime() time.Time {
	ts, err := cc.CCAPI.GetTxTimestamp()
	PanicError(err)
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
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

// This call is only supported in a read only transaction.
func (cc CommonChaincode) GetStateByRangeWithPagination(startKey, endKey string, pageSize int, bookmark string) (shim.StateQueryIteratorInterface, QueryResponseMetadata) {
	var iteratorInterface, r, err = cc.CCAPI.GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	PanicError(err)
	return iteratorInterface, QueryResponseMetadata{int(r.FetchedRecordsCount), r.Bookmark}
}

func (cc CommonChaincode) SetEvent(name string, payload []byte) {
	var err = cc.CCAPI.SetEvent(name, payload)
	PanicError(err)
}

var DeferHandlerPeerResponse = func(errString string, params ...interface{}) bool {
	var response = params[0].(*peer.Response)
	response.Status = shim.ERROR
	response.Message = errString
	fmt.Println("DeferHandlerPeerResponse", errString)
	debug.PrintStack()
	return true
}

func (cc CommonChaincode) GetFunctionAndArgs() (string, [][]byte) {
	var allArgs = cc.CCAPI.GetArgs()
	var fcn = ""
	var args = [][]byte{}
	if len(allArgs) >= 1 {
		fcn = string(allArgs[0])
		args = allArgs[1:]
	}
	return fcn, args
}
func ChaincodeStart(cc shim.Chaincode) {
	var err = shim.Start(cc)
	PanicError(err)
}
