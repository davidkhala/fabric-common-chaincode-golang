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
	MspID          string
	Cert           *x509.Certificate `json:"-"`
	CertificatePem []byte
	Attrs          attrmgr.Attributes
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
	block, rest := pem.Decode(idbytes)
	AssertEmpty(rest, "pem decode failed:"+string(rest))
	c.CertificatePem = block.Bytes
	c.Cert = c.GetCert()
	attrs, err := attrmgr.New().GetAttributesFromCert(c.Cert)
	PanicError(err)
	c.Attrs = *attrs
	return c
}
func (c ClientIdentity) GetCert() *x509.Certificate {
	cert, err := x509.ParseCertificate(c.CertificatePem)
	PanicError(err)
	return cert
}

func (c *ClientIdentity) GetAttributeValue(attrName string) (string) {
	return c.Attrs.Attrs[attrName]
}
