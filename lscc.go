package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/protos/common"
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

type ChaincodeData struct {
	Name                string                         // Name of the chaincode
	Version             string                         // Version of the chaincode
	Escc                string                         // Escc for the chaincode instance
	Vscc                string                         // Vscc for the chaincode instance
	Data                ccprovider.CDSData             // Data data specific to the package
	Policy              common.SignaturePolicyEnvelope // Policy endorsement policy for the chaincode instance
	InstantiationPolicy common.SignaturePolicyEnvelope // InstantiationPolicy for the chaincode
}

//type SignaturePolicyEnvelope struct {
//	Version    int32
//	Rule       *common.SignaturePolicy
//	Identities []*msp.MSPPrincipal
//}

func (cc CommonChaincode) GetChaincodeData(channel, checkedChaincode string) ChaincodeData {
	var args = [][]byte{[]byte("GetChaincodeData"), []byte(channel), []byte(checkedChaincode)}
	var resp = cc.InvokeChaincode("lscc", args, operationChannel)

	var chaincodeData = ccprovider.ChaincodeData{}
	var err = proto.Unmarshal(resp.Payload, &chaincodeData)
	PanicError(err)

	var policyProto = common.SignaturePolicyEnvelope{}
	PanicError(proto.Unmarshal(chaincodeData.Policy, &policyProto)) //TODO TBC  [principal:"\n\010ASTRIMSP"  principal:"\n\007icddMSP" ]

	var instantiatePolicyProto = common.SignaturePolicyEnvelope{}
	PanicError(proto.Unmarshal(chaincodeData.InstantiationPolicy, &instantiatePolicyProto))

	var convertedChaincodeData = ChaincodeData{
		Name:                chaincodeData.Name,
		Version:             chaincodeData.Version,
		Escc:                chaincodeData.Escc,
		Vscc:                chaincodeData.Vscc,
		Policy:              policyProto,
		InstantiationPolicy: instantiatePolicyProto,
	}
	return convertedChaincodeData
}
