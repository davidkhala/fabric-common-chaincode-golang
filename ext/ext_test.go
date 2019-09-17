package ext

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
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
	var certPem = `-----BEGIN CERTIFICATE-----
MIICEzCCAbqgAwIBAgIQQOJKqT3HF9Jz8qM56o2otjAKBggqhkjOPQQDAjBwMQsw
CQYDVQQGEwJDTjERMA8GA1UECBMIU2hhbmdoYWkxETAPBgNVBAcTCFNoYW5naGFp
MRswGQYDVQQKExJ0ZXN0LmNhbGxkb2N0b3IuaGsxHjAcBgNVBAMTFWNhLnRlc3Qu
Y2FsbGRvY3Rvci5oazAeFw0xOTAyMDMwNzIzNDRaFw0yOTAxMzEwNzIzNDRaMGkx
CzAJBgNVBAYTAkNOMREwDwYDVQQIEwhTaGFuZ2hhaTERMA8GA1UEBxMIU2hhbmdo
YWkxDzANBgNVBAsTBmNsaWVudDEjMCEGA1UEAwwaVXNlckFwcEB0ZXN0LmNhbGxk
b2N0b3IuaGswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASVHJJ/intU0Kbe3iHn
ugQlGbLRVJDM90y3Jx64RHJPvNiICvtQ8srN0OylkD+fjXVpfGH+BzF6cBRca7mf
EPNFoz0wOzAMBgNVHRMBAf8EAjAAMCsGA1UdIwQkMCKAIIPRV0eT91+L9peQ6/IG
jPAYvlSRsJUS87PlwA5LqIm3MAoGCCqGSM49BAMCA0cAMEQCIGcBMwNGZ0Gyx5Eh
/vbSdTYhn6r6klP82V3qS3bVgU1WAiBh/rDpaNTcvgSh2rrP6GuLcfAm9js51gVJ
LRu3mTbvVQ==
-----END CERTIFICATE-----`
	var token = "PLMLVSKWSU95"
	var signatureBytes = []byte{
		48, 68, 2, 32, 118, 240, 127, 79, 161, 89, 125, 155, 122, 131, 29, 91, 226, 42, 163, 58, 26, 87, 195, 61, 49, 237, 251, 46, 42, 216, 105, 188, 170, 30, 149, 225, 2, 32, 43, 104, 81, 138, 133, 142, 84, 176, 205, 230, 38, 139, 114, 118, 15, 134, 165, 19, 176, 205, 251, 10, 190, 160, 192, 223, 141, 181, 235, 181, 57, 97,
	}

	var signID = "BC"
	var certBytes = []byte(certPem)

	sm := &entities.SignedMessage{
		ID:      []byte(signID),
		Payload: []byte(token),
		Sig:     signatureBytes,
	}
	var valid = VerifyECDSA(sm, certBytes)
	t.Log("valid", valid)

}
