package golang

import "github.com/hyperledger/fabric/core/chaincode/shim"

func (cc *CommonChaincode) NewMock() *shim.MockStub {
	var mock = shim.NewMockStub(cc.Name, cc)
	cc.Prepare(mock)
	return mock
}
