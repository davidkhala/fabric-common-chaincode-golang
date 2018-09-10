package golang

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"testing"
	"github.com/hyperledger/fabric/common/ledger/testutil"
	. "github.com/davidkhala/goutils"
)

type TestChaincode struct {
	CommonChaincode
}

const (
	name = "TestChaincode"
)

var logger = shim.NewLogger(name)

func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("########### " + name + " Init ###########")
	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *TestChaincode) Invoke(ccAPI shim.ChaincodeStubInterface) peer.Response {
	logger.Info("########### " + name + " Invoke ###########")
	fcn, _ := ccAPI.GetFunctionAndParameters()
	switch fcn {
	case "panic":
		PanicString("use panic")
	}
	return shim.Success([]byte(nil))
}

var cc = new(TestChaincode)
var mock = shim.NewMockStub(name, cc)
//initialize mocker
func TestCommonChaincode_Prepare(t *testing.T) {
	cc.Prepare(mock)
}

func TestTestChaincode_Init(t *testing.T) {
	var initArgs [][]byte
	initArgs = append(initArgs, []byte("Initfcn")) //fcn
	var TxID = "ob"

	var response = mock.MockInit(TxID, initArgs)
	t.Log("init ", response)
	testutil.AssertSame(t, response.Status, int32(200));
}
func TestTestChaincode_Invoke(t *testing.T) {

	var args [][]byte
	args = append(args, []byte("Invokefcn")) //fcn

	var TxID = "oa"
	var response = mock.MockInvoke(TxID, args)
	t.Log("invoke ", response)
	testutil.AssertSame(t, response.Status, int32(200));
	//	when error status is 500
}

func TestCreateCompositeKey(t *testing.T) {
	var cKey1 = cc.CreateCompositeKey("a", []string{"c", "C"})
	var cKey2 = cc.CreateCompositeKey("a", []string{"d", "D"})
	var TxID = "composityKey"
	mock.MockTransactionStart(TxID);
	cc.PutState(cKey1, []byte("c"))
	cc.PutState(cKey2, []byte("C"))
	mock.MockTransactionEnd(TxID);
	TxID = "composite1"
	mock.MockTransactionStart(TxID)
	iterator := cc.GetStateByPartialCompositeKey("a", []string{"d"})
	var kvs States
	kvs.ParseStates(iterator)
	t.Log(kvs)
	mock.MockTransactionEnd(TxID)
}

/**
[31m2018-07-09 12:46:27.277 HKT [mock] HasNext -> ERRO 001[0m HasNext() couldn't get Current
mockstub.go line 410:	mockLogger.Error("HasNext() couldn't get Current")
 */
func TestWorldStates(t *testing.T) {
	var TxID = "composityKey"
	mock.MockTransactionStart(TxID)

	cc.PutState("a_1", []byte("c"))
	cc.PutState("a_2", []byte("C"))
	cc.PutState("a_3", []byte("C"))

	mock.MockTransactionEnd(TxID);
	TxID = "composite1"
	mock.MockTransactionStart(TxID)
	kvs := cc.WorldStates( "")

	t.Log(kvs)

	kvs.ParseStates(cc.GetStateByRange("a_1", ""))
	t.Log(kvs)
	mock.MockTransactionEnd(TxID)
}

func TestGetStateObj(t *testing.T) {

	var value = KVJson{"1", "2", "3"}
	var TxID = "compositeKey"
	var key = "a_1"
	mock.MockTransactionStart(TxID)

	cc.PutStateObj(key, value)

	mock.MockTransactionEnd(TxID);
	TxID = "composite1"
	mock.MockTransactionStart(TxID)

	var recovered KVJson
	cc.GetStateObj(key, &recovered)

	t.Log(recovered)
	mock.MockTransactionEnd(TxID)
}
func TestModifyValue(t *testing.T) {
	var key = "a_1"
	var kv KVJson
	TxID := "composite2"
	mock.MockTransactionStart(TxID)
	var modifier = func(v interface{}) {
		t.Log("modifierTest", v)
		t.Log("modifierTest", kv)
	}

	cc.ModifyValue(key, modifier, &kv)
	mock.MockTransactionEnd(TxID)
}
