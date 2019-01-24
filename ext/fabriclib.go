package ext

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"encoding/pem"
	"crypto/x509"
	"github.com/hyperledger/fabric/bccsp/factory"
	"fmt"
	"github.com/hyperledger/fabric/bccsp/utils"
)

//Note:clientIdentityImpl has no public properties, so ToJson(cid.ClientIdentity) is empty
func NewClientIdentity(stub shim.ChaincodeStubInterface) cid.ClientIdentity {
	var identity, err = cid.New(stub)
	PanicError(err)
	return identity
}

type KeyEndorsementPolicy struct {
	statebased.KeyEndorsementPolicy
}

func NewKeyEndorsementPolicy(clonedPolicy []byte) KeyEndorsementPolicy {
	var result, err = statebased.NewStateEP(clonedPolicy)
	PanicError(err)
	return KeyEndorsementPolicy{result}
}
func (t KeyEndorsementPolicy) Policy() []byte {
	var result, err = t.KeyEndorsementPolicy.Policy()
	PanicError(err)
	return result
}

func (t KeyEndorsementPolicy) AddOrgs(roleType msp.MSPRole_MSPRoleType, organizations ...string) {
	var err = t.KeyEndorsementPolicy.AddOrgs(statebased.RoleType(roleType.String()), organizations...)
	PanicError(err)
}

func Sign(ID string, payload []byte, signKeyPemBytes []byte) (*entities.SignedMessage, error) {
	msg := entities.SignedMessage{
		ID: []byte(ID),
		Payload: payload,
	}
	factory.InitFactories(nil)
	entPvt, err := entities.NewECDSASignerEntity(string(msg.ID), factory.GetDefault(), signKeyPemBytes)
	if err !=nil {
		return nil, fmt.Errorf("failed to new signer entity, err: %s", err)
	}

	err = msg.Sign(entPvt)
	if err !=nil {
		return nil, fmt.Errorf("failed to sign, err: %s", err)
	}
	return &msg, nil
}

func Verify(msg *entities.SignedMessage, signCertBytes []byte) (bool, error) {
	pemCert, _ := pem.Decode([]byte(signCertBytes))

	// get a cert
	cert, err := x509.ParseCertificate(pemCert.Bytes)
	if err !=nil {
		return false, fmt.Errorf("failed to parse certificate, err: %s", err)
	}

	publicKeyPemBytes, err := utils.PublicKeyToPEM(cert.PublicKey, nil)
	if err !=nil {
		return false, fmt.Errorf("failed to convert pub key to pem, err: %s", err)
	}
	factory.InitFactories(nil)
	entPub, err := entities.NewECDSAVerifierEntity(string(msg.ID), factory.GetDefault(), publicKeyPemBytes)
	if err !=nil {
		return false, fmt.Errorf("failed to new verifier entity, err: %s", err)
	}

	valid, err := msg.Verify(entPub)
	if err !=nil {
		return false, fmt.Errorf("failed to verify, err: %s", err)
	}
	return valid, nil
}