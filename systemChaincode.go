package golang

import "strings"

// TODO the meaning of ChaincodeExist
func (cc CommonChaincode) ChaincodeExist(operationChannel, channel, checkedChaincode string) bool {
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
