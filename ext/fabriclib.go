package ext

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
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

func (t KeyEndorsementPolicy) AddOrgs(roleType msp.MSPRole_MSPRoleType, organizations ...string) {
	var err = t.KeyEndorsementPolicy.AddOrgs(statebased.RoleType(roleType.String()), organizations...)
	PanicError(err)
}
