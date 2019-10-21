package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/common/ccprovider"
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
func (cc CommonChaincode) GetChaincodeData(channel, checkedChaincode string) ccprovider.ChaincodeData {
	var args = [][]byte{[]byte("GetChaincodeData"), []byte(channel), []byte(checkedChaincode)}
	var resp = cc.InvokeChaincode("lscc", args, operationChannel)

	var chaincodeData = ccprovider.ChaincodeData{}
	var err = proto.Unmarshal(resp.Payload, &chaincodeData)
	PanicError(err)
	return chaincodeData
}
