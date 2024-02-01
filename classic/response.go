package classic

import (
	"fmt"
	"github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"runtime/debug"
)

func PanicPeerResponse(resp peer.Response) {
	if resp.Status >= shim.ERRORTHRESHOLD {
		var errorPB = PeerResponse{
			resp.Status,
			resp.Message,
			string(resp.Payload),
		}
		goutils.PanicString(string(goutils.ToJson(errorPB)))
	}
}

var DeferHandlerPeerResponse = func(errString string, params ...interface{}) bool {
	var response = params[0].(*peer.Response)
	response.Status = shim.ERROR
	response.Message = errString
	response.Payload = []byte(errString)
	fmt.Println("DeferHandlerPeerResponse", errString)
	debug.PrintStack()
	return true
}

// PeerResponse a readable structure of peer.response
type PeerResponse struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `json:"status,omitempty"`
	// A message associated with the response code.
	Message string `json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload string `json:"payload,omitempty"`
}
