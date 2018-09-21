package golang

import (
	"crypto/x509"
	"encoding/pem"
	. "github.com/davidkhala/goutils"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/attrmgr"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

// alternative of creator starting from 1.1
type ClientIdentity struct {
	MspID string
	Cert  x509.Certificate
	Attrs attrmgr.Attributes
}

func NewClientIdentity(stub shim.ChaincodeStubInterface) (c ClientIdentity) {
	signingID := &msp.SerializedIdentity{}
	creator, err := stub.GetCreator()
	PanicError(err)
	if creator == nil {
		panic(errors.New("failed to get transaction invoker's identity from the chaincode stub"))
	}
	err = proto.Unmarshal(creator, signingID)
	PanicError(err)

	c.MspID = signingID.GetMspid()
	idbytes := signingID.GetIdBytes()
	block, _ := pem.Decode(idbytes)
	if block == nil {
		panic(errors.New("Expecting a PEM-encoded X509 certificate; PEM block not found"))
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	PanicError(err)
	c.Cert = *cert
	attrs, err := attrmgr.New().GetAttributesFromCert(cert)
	PanicError(err)
	c.Attrs = *attrs
	return c
}

func (c *ClientIdentity) GetAttributeValue(attrName string) (string) {
	return c.Attrs.Attrs[attrName]
}