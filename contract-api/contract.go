package contract_api

import (
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

type ExtendedContract contractapi.Contract

func (c *ExtendedContract) GetIgnoredFunctions() []string {
	return []string{
		"GetClientIdentity", "GetStub",
	}
}
func (c *ExtendedContract) GetStub() shim.ChaincodeStubInterface {
	var context = c.TransactionContextHandler.(*contractapi.TransactionContext)
	return context.GetStub()
}
func (c *ExtendedContract) GetClientIdentity() cid.ClientIdentity {
	var context = c.TransactionContextHandler.(*contractapi.TransactionContext)
	return context.GetClientIdentity()
}

func (c *ExtendedContract) GetInfo() metadata.InfoMetadata {
	return c.Info
}

func (c *ExtendedContract) GetUnknownTransaction() interface{} {
	return c.UnknownTransaction
}

func (c *ExtendedContract) GetBeforeTransaction() interface{} {
	return c.BeforeTransaction
}

func (c *ExtendedContract) GetAfterTransaction() interface{} {
	return c.AfterTransaction
}

// GetName returns the name of the contract
func (c *ExtendedContract) GetName() string {
	return c.Name
}

// GetTransactionContextHandler returns the current transaction context set for
// the contract. If none has been set then TransactionContext will be returned
func (c *ExtendedContract) GetTransactionContextHandler() contractapi.SettableTransactionContextInterface {
	if c.TransactionContextHandler == nil {
		return new(contractapi.TransactionContext)
	}

	return c.TransactionContextHandler
}
