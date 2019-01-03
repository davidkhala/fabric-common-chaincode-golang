package golang

import . "github.com/davidkhala/goutils"

func (t CommonChaincode) SetStateValidationParameter(key string, endorsementPolicy []byte) {
	var err = t.CCAPI.SetStateValidationParameter(key, endorsementPolicy)
	PanicError(err)
}
func (t CommonChaincode) GetStateValidationParameter(key string) []byte {
	var result, err = t.CCAPI.GetStateValidationParameter(key)
	PanicError(err)
	return result
}
func (t CommonChaincode) SetPrivateDataValidationParameter(collection, key string, endorsementPolicy []byte) {
	var err = t.CCAPI.SetPrivateDataValidationParameter(collection, key, endorsementPolicy)
	PanicError(err)
}
func (t CommonChaincode) GetPrivateDataValidationParameter(collection, key string) []byte {
	var result, err = t.CCAPI.GetPrivateDataValidationParameter(collection, key)
	PanicError(err)
	return result
}
