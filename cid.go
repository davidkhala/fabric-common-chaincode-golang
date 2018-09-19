package golang

import (
	"crypto/x509"
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//import "github.com/hyperledger/fabric/core/chaincode/lib/cid"
// alternative of creator starting from 1.1
type ClientIdentity struct {
	cid cid.ClientIdentity
}

func NewClientIdentity(stub shim.ChaincodeStubInterface) ClientIdentity {
	var c, err = cid.New(stub)
	PanicError(err)
	return ClientIdentity{c}
}

func (t ClientIdentity) GetID() string {
	var id, err = t.cid.GetID()
	PanicError(err)
	return id
}
func (t ClientIdentity) GetMSPID() string {
	var mspid, err = t.cid.GetMSPID()
	PanicError(err)
	return mspid
}
func (t ClientIdentity) GetX509Certificate() *x509.Certificate {
	var cert, err = t.cid.GetX509Certificate()
	PanicError(err)
	return cert
}
func (t ClientIdentity) GetAttributeValue(attrName string) (string, bool) {
	var value, found, err = t.cid.GetAttributeValue(attrName)
	PanicError(err)
	return value, found
}
func (t ClientIdentity) AssertAttributeValue(attrName, attrValue string) {
	var err = t.cid.AssertAttributeValue(attrName, attrValue)
	PanicError(err)
}
