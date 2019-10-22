package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/msp"
	"strings"
)

const operationChannel = "" // It does not matter what channel we use, so use current channel

// query function, a short representation of GetChaincodeData
func (cc CommonChaincode) ChaincodeExist(channel, checkedChaincode string) bool {
	var args = [][]byte{[]byte("ChaincodeExists"), []byte(channel), []byte(checkedChaincode)}
	var resp = cc.CCAPI.InvokeChaincode("lscc", args, operationChannel)
	//  {"status":500,"message":"could not find chaincode with name 'any'"
	if resp.Status == 500 && strings.Contains(resp.Message, "could not find chaincode with name") {
		return false
	} else {
		PanicPeerResponse(resp)
	}
	return true

}

type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification msp.MSPPrincipal_Classification
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	// For the Combined Classification type, the Principal is a marshalled
	// CombinedPrincipal.
	Principal msp.MSPRole
}

type SignaturePolicyEnvelope struct {
	Version    int32
	Rule       *common.SignaturePolicy // it has to be a pointer to preserve raw type: SignaturePolicy_NOutOf_|SignaturePolicy_SignedBy
	Identities []MSPPrincipal
}

func (t *SignaturePolicyEnvelope) LoadIdentities(identities []*msp.MSPPrincipal) {
	var result []MSPPrincipal
	for _, mspPrincipal := range identities {
		var principal = &msp.MSPRole{}
		PanicError(proto.Unmarshal(mspPrincipal.Principal, principal))
		result = append(result, MSPPrincipal{
			PrincipalClassification: mspPrincipal.PrincipalClassification,
			Principal:               *principal,
		})
	}
	t.Identities = result
}

type ChaincodeData struct {
	Name                string                  // Name of the chaincode
	Version             string                  // Version of the chaincode
	Escc                string                  // Escc for the chaincode instance
	Vscc                string                  // Vscc for the chaincode instance
	Data                string                  // Data data specific to the package
	Policy              SignaturePolicyEnvelope // Policy endorsement policy for the chaincode instance
	InstantiationPolicy SignaturePolicyEnvelope // InstantiationPolicy for the chaincode
}

func (cc CommonChaincode) GetChaincodeData(channel, checkedChaincode string) ChaincodeData {
	var args = [][]byte{[]byte("GetChaincodeData"), []byte(channel), []byte(checkedChaincode)}
	var resp = cc.InvokeChaincode("lscc", args, operationChannel)

	var chaincodeData = ccprovider.ChaincodeData{}
	var err = proto.Unmarshal(resp.Payload, &chaincodeData)
	PanicError(err)

	var policyProto common.SignaturePolicyEnvelope
	PanicError(proto.Unmarshal(chaincodeData.Policy, &policyProto))
	var policyText = SignaturePolicyEnvelope{
		Version: policyProto.Version,
		Rule:    policyProto.Rule,
	}
	policyText.LoadIdentities(policyProto.Identities)

	var instantiatePolicyProto common.SignaturePolicyEnvelope
	PanicError(proto.Unmarshal(chaincodeData.InstantiationPolicy, &instantiatePolicyProto))
	var instantiatePolicyText = SignaturePolicyEnvelope{
		Version: instantiatePolicyProto.Version,
		Rule:    instantiatePolicyProto.Rule,
	}
	instantiatePolicyText.LoadIdentities(instantiatePolicyProto.Identities)

	var dataProto ccprovider.CDSData
	PanicError(proto.Unmarshal(chaincodeData.Data, &dataProto))
	var convertedChaincodeData = ChaincodeData{
		Name:                chaincodeData.Name,
		Version:             chaincodeData.Version,
		Escc:                chaincodeData.Escc,
		Vscc:                chaincodeData.Vscc,
		Data:                string(ToJson(dataProto)),
		Policy:              policyText,
		InstantiationPolicy: instantiatePolicyText,
	}
	return convertedChaincodeData
}
