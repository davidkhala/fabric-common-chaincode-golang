package ext

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
)

func ParseClientIdentity(stub shim.ChaincodeStubInterface) cid.ClientIdentity {
	var identity, err = cid.New(stub)
	PanicError(err)
	return identity
}
