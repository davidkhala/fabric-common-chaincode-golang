package golang

import . "github.com/davidkhala/goutils"

func (cc CommonChaincode) SetStateValidationParameter(key string, endorsementPolicy []byte) {
	var err = cc.CCAPI.SetStateValidationParameter(key, endorsementPolicy)
	PanicError(err)
}
func (cc CommonChaincode) GetStateValidationParameter(key string) []byte {
	var result, err = cc.CCAPI.GetStateValidationParameter(key)
	PanicError(err)
	return result
}
func (cc CommonChaincode) SetPrivateDataValidationParameter(collection, key string, endorsementPolicy []byte) {
	var err = cc.CCAPI.SetPrivateDataValidationParameter(collection, key, endorsementPolicy)
	PanicError(err)
}
func (cc CommonChaincode) GetPrivateDataValidationParameter(collection, key string) []byte {
	var result, err = cc.CCAPI.GetPrivateDataValidationParameter(collection, key)
	PanicError(err)
	return result
}
