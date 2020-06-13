package contract_api

import (
	"github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewChaincode(contracts ...contractapi.ContractInterface) *contractapi.ContractChaincode {
	cc, err := contractapi.NewChaincode(contracts...)
	goutils.PanicError(err)
	return cc
}
