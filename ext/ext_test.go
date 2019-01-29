package ext

import (
	"fmt"
	"github.com/hyperledger/fabric/protos/msp"
	"testing"
)

//pointer test
func TestNewKeyEndorsementPolicy(t *testing.T) {
	var keyPolicy = NewKeyEndorsementPolicy(nil)
	fmt.Println("empty policy", string(keyPolicy.Policy()))
	keyPolicy.AddOrgs(msp.MSPRole_MEMBER, "org1MSP")
	fmt.Println("non-empty policy", string(keyPolicy.Policy()))
}
