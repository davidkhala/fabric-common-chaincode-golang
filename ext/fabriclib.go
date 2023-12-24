package ext

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go-apiv2/msp"
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
)

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

// Success ...
func Success(payload []byte) peer.Response {
	return peer.Response{
		Status:  shim.OK,
		Payload: payload,
	}
}

// Error ...
func Error(msg string) peer.Response {
	return peer.Response{
		Status:  shim.ERROR,
		Message: msg,
	}
}
