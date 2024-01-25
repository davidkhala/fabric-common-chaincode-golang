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
