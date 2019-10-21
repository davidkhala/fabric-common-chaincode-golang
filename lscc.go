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

type SignaturePolicyEnvelope struct {
	Version    int32
	Rule       common.SignaturePolicy
	Identities []*msp.MSPPrincipal
}
type ChaincodeData struct {
	Name                string                  // Name of the chaincode
	Version             string                  // Version of the chaincode
	Escc                string                  // Escc for the chaincode instance
	Vscc                string                  // Vscc for the chaincode instance
	Data                []byte `json:"data"`    // Data data specific to the package
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
	PanicError(proto.Unmarshal(chaincodeData.Policy, &policyProto)) //TODO TBC  [principal:"\n\010ASTRIMSP"  principal:"\n\007icddMSP" ]
	var policyText = SignaturePolicyEnvelope{
		Version:    policyProto.Version,
		Rule:       *policyProto.Rule,
		Identities: policyProto.Identities,
	}

	var instantiatePolicyProto common.SignaturePolicyEnvelope
	PanicError(proto.Unmarshal(chaincodeData.InstantiationPolicy, &instantiatePolicyProto))
	var instantiatePolicyText = SignaturePolicyEnvelope{
		Version:    instantiatePolicyProto.Version,
		Rule:       *instantiatePolicyProto.Rule,
		Identities: instantiatePolicyProto.Identities,
	}

	//TODO MSPRole
	var dataProto ccprovider.CDSData
	PanicError(proto.Unmarshal(chaincodeData.Data, &dataProto))
	var convertedChaincodeData = ChaincodeData{
		Name:                chaincodeData.Name,
		Version:             chaincodeData.Version,
		Escc:                chaincodeData.Escc,
		Vscc:                chaincodeData.Vscc,
		Data:                ToJson(dataProto), //TODO to test bytes
		Policy:              policyText,
		InstantiationPolicy: instantiatePolicyText,
	}
	return convertedChaincodeData
}
