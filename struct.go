package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"time"
)

type KeyModification struct {
	TxId      string
	Value     []byte
	Timestamp time.Time
	IsDelete  bool
}

func ParseHistory(iterator shim.HistoryQueryIteratorInterface, filter func(KeyModification) bool) []KeyModification {
	defer PanicError(iterator.Close())
	var result []KeyModification
	for iterator.HasNext() {
		keyModification, err := iterator.Next()
		PanicError(err)
		var timeStamp = keyModification.Timestamp

		var translated = KeyModification{
			keyModification.TxId,
			keyModification.Value,
			timeStamp.AsTime(),
			keyModification.IsDelete}
		if filter == nil || filter(translated) {
			result = append(result, translated)
		}
	}
	return result
}

type StateKV struct {
	Namespace string
	Key       string
	Value     string
}
type QueryResponseMetadata struct {
	FetchedRecordsCount int
	Bookmark            string
}

func ParseStates(iterator shim.StateQueryIteratorInterface, filter func(StateKV) bool) []StateKV {
	defer PanicError(iterator.Close())
	var kvs []StateKV
	for iterator.HasNext() {
		kv, err := iterator.Next()
		PanicError(err)
		var stateKV = StateKV{kv.Namespace, kv.Key, string(kv.Value)}
		if filter == nil || filter(stateKV) {
			kvs = append(kvs, stateKV)
		}
	}
	return kvs
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
