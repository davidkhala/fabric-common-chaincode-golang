package contract_api

import (
	"testing"
)

type StupidContract struct {
	ExtendedContract
}

func (c *StupidContract) Ping() {

}
func TestBuild(t *testing.T) {

	Start(NewChaincode(&StupidContract{}))

}
