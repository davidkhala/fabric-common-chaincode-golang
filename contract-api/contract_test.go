package contract_api

import (
	"errors"
	"github.com/davidkhala/goutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StupidContract struct {
	ExtendedContract
}

func (*StupidContract) Defer() (err error) {
	defer goutils.Deferred(DeferHandlerError, &err)
	panic(errors.New("defer"))
}

func (c *StupidContract) Ping() {

}
func TestBuild(t *testing.T) {

	Start(NewChaincode(&StupidContract{}))

}
func TestDefer(t *testing.T) {
	var contract = StupidContract{}
	err := contract.Defer()
	assert.EqualError(t, err, "defer")
}
