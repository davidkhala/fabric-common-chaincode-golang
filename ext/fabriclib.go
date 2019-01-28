package ext

import (
	. "github.com/davidkhala/goutils"
	"github.com/davidkhala/goutils/crypto"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased"
	"github.com/hyperledger/fabric/protos/msp"
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

func (t KeyEndorsementPolicy) AddOrgs(roleType msp.MSPRole_MSPRoleType, MSPIDs ...string) {
	var err = t.KeyEndorsementPolicy.AddOrgs(statebased.RoleType(roleType.String()), MSPIDs...)
	PanicError(err)
}

func NewECDSASignerEntity(ID string, signKeyBytes []byte) *entities.BCCSPSignerEntity {
	factory.InitFactories(nil)
	signerEntity, err := entities.NewECDSASignerEntity(ID, factory.GetDefault(), signKeyBytes)
	PanicError(err)
	return signerEntity
}

func Sign(msg *entities.SignedMessage, signer entities.Signer) {
	err := msg.Sign(signer) // msg sig changed inline
	PanicError(err)
}

func SignECDSA(msg *entities.SignedMessage, privateKeyBytes []byte) {
	var privateSigner = NewECDSASignerEntity(string(msg.ID), privateKeyBytes)
	Sign(msg, privateSigner)
}
func NewECDSAVerifierEntity(ID string, publicKeyBytes []byte) *entities.BCCSPSignerEntity {
	factory.InitFactories(nil)
	pub, err := entities.NewECDSAVerifierEntity(ID, factory.GetDefault(), publicKeyBytes)
	PanicError(err)
	return pub
}

func Verify(msg *entities.SignedMessage, signer entities.Signer) bool {
	valid, err := msg.Verify(signer)
	PanicError(err)
	return valid
}
func VerifyECDSA(msg *entities.SignedMessage, certBytes []byte) bool {
	var cert = crypto.ParseCertPem(certBytes)
	publicKeyPemBytes, err := utils.PublicKeyToPEM(cert.PublicKey, nil)
	PanicError(err)
	var pub = NewECDSAVerifierEntity(string(msg.ID), publicKeyPemBytes)
	return Verify(msg, pub)
}
