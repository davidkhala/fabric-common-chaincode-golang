package ext

import (
	"fmt"
	"github.com/hyperledger/fabric/protos/msp"
	"testing"
)

// pointer test
func TestNewKeyEndorsementPolicy(t *testing.T) {
	var keyPolicy = NewKeyEndorsementPolicy(nil)
	fmt.Println("empty policy", string(keyPolicy.Policy()))
	keyPolicy.AddOrgs(msp.MSPRole_MEMBER, "org1MSP")
	fmt.Println("non-empty policy", string(keyPolicy.Policy()))
}
func TestVerifyECDSA(t *testing.T) {
	var signatureBytes = []byte{
		48, 68, 2, 32, 118, 240, 127, 79, 161, 89, 125, 155, 122, 131, 29, 91, 226, 42, 163, 58, 26, 87, 195, 61, 49, 237, 251, 46, 42, 216, 105, 188, 170, 30, 149, 225, 2, 32, 43, 104, 81, 138, 133, 142, 84, 176, 205, 230, 38, 139, 114, 118, 15, 134, 165, 19, 176, 205, 251, 10, 190, 160, 192, 223, 141, 181, 235, 181, 57, 97
	}
}
