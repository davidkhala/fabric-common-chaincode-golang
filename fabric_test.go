package golang

import (
	"github.com/davidkhala/fabric-common-chaincode-golang/cid"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestChaincode struct {
	CommonChaincode
}

const (
	name = "TestChaincode"
)

func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	t.Prepare(stub)
	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) (response peer.Response) {
	defer Deferred(DeferHandlerPeerResponse, &response)
	var fcn, _ = stub.GetFunctionAndParameters()
	var responseBytes []byte
	switch fcn {
	case "panic":
		PanicString("use panic")
	case "whoami":
		responseBytes = ToJson(cid.NewClientIdentity(stub))
	}
	return shim.Success(responseBytes)
}

var cc = new(TestChaincode)
var mock = shim.NewMockStub(name, cc)

//initialize mocker
func TestCommonChaincode_Prepare(t *testing.T) {
	cc.Prepare(mock)
}

func TestTestChaincode_Init(t *testing.T) {
	var args = ArgsBuilder("Initfcn")
	var TxID = "ob"

	var response = mock.MockInit(TxID, args.Get())
	t.Log("init ", response)
	assert.Equal(t, int32(200), response.Status)
}
func TestTestChaincode_Invoke(t *testing.T) {

	var args = ArgsBuilder("Invokefcn")

	var TxID = "oa"
	var response = mock.MockInvoke(TxID, args.Get())
	t.Log("invoke ", response)
	assert.Equal(t, int32(200), response.Status)
	//	when error status is 500
}

func TestCreateCompositeKey(t *testing.T) {
	var cKey1 = cc.CreateCompositeKey("a", []string{"c", "C"})
	var cKey2 = cc.CreateCompositeKey("a", []string{"d", "D"})
	var TxID = "composityKey"
	mock.MockTransactionStart(TxID)
	cc.PutState(cKey1, []byte("c"))
	cc.PutState(cKey2, []byte("C"))
	mock.MockTransactionEnd(TxID)
	TxID = "composite1"
	mock.MockTransactionStart(TxID)
	iterator := cc.GetStateByPartialCompositeKey("a", []string{"d"})
	var kvs = ParseStates(iterator, nil)
	t.Log(kvs)
	mock.MockTransactionEnd(TxID)
}

func TestWorldStates(t *testing.T) {
	var TxID = "composityKey"
	mock.MockTransactionStart(TxID)

	cc.PutState("a_1", []byte("c"))
	cc.PutState("a_2", []byte("C"))
	cc.PutState("a_3", []byte("C"))

	mock.MockTransactionEnd(TxID)
	TxID = "composite1"
	mock.MockTransactionStart(TxID)
	kvs := cc.WorldStates("", nil)

	t.Log(kvs)

	var states2 = ParseStates(cc.GetStateByRange("a_1", ""), nil)
	t.Log(states2)
	mock.MockTransactionEnd(TxID)
}

func TestGetStateObj(t *testing.T) {

	var value = StateKV{"1", "2", "3"}
	var TxID = "compositeKey"
	var key = "a_1"
	mock.MockTransactionStart(TxID)

	cc.PutStateObj(key, value)

	mock.MockTransactionEnd(TxID)
	TxID = "composite1"
	mock.MockTransactionStart(TxID)

	var recovered StateKV
	cc.GetStateObj(key, &recovered)

	t.Log(recovered)
	mock.MockTransactionEnd(TxID)
}
func TestTouch(t *testing.T) {
	touch()
}
