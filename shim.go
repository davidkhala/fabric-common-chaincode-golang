package golang

import (
	"github.com/davidkhala/fabric-common-chaincode-golang/cid"
	"github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

// GetMSPID From https://github.com/hyperledger/fabric-chaincode-go/commit/2d899240a7ed642a381ba9df2f6b0c303cb149dc
func GetMSPID() cid.MSPID {
	var mspId, err = shim.GetMSPID()
	goutils.PanicError(err)
	return mspId
}
func ChaincodeStart(cc shim.Chaincode) {
	var err = shim.Start(cc)
	goutils.PanicError(err)
}
