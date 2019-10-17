package golang

import "strings"

func (cc CommonChaincode) ChaincodeExist(channel, checkedChaincode string) bool {
	var args = [][]byte{[]byte("ChaincodeExists"), []byte(channel), []byte(checkedChaincode)}
	var resp = cc.InvokeChaincode("lscc", args, channel) // TODO could we specify another channel?
	//  {"status":500,"message":"could not find chaincode with name 'any'"
	if resp.Status == 500 && strings.Contains(resp.Message, "could not find chaincode with name") {
		return false
	} else {
		PanicPeerResponse(resp)
	}
	return true

}
