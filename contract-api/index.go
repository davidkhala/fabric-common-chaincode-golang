package contract_api

import (
	"github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Start(chaincode *contractapi.ContractChaincode) {
	var err = chaincode.Start()
	switch err.Error() {
	case "'CORE_CHAINCODE_ID_NAME' must be set":
		println("[dev mode]", err.Error())
		break
	default:
		goutils.PanicError(err)
	}

}
func NewChaincode(contracts ...contractapi.ContractInterface) *contractapi.ContractChaincode {
	chaincode, err := contractapi.NewChaincode(contracts...)
	goutils.PanicError(err)
	return chaincode
}

type DeferHandler func(error) (success bool)

func Deferred(handler DeferHandler) {
	err := recover()
	if err == nil {
		return
	}

	var success = handler(err.(error))
	if !success {
		panic(err)
	}
}
func DefaultDeferHandler(err *error) func(error) bool {
	return func(_err error) bool {
		*err = _err
		return true
	}
}
