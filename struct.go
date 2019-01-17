package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type KeyModification struct {
	TxId      string
	Value     []byte
	Timestamp TimeLong
	IsDelete  bool
}

func ParseHistory(iterator shim.HistoryQueryIteratorInterface, filter Filter) []KeyModification {
	defer iterator.Close()
	var result []KeyModification
	for iterator.HasNext() {
		keyModification, err := iterator.Next()
		PanicError(err)
		var timeStamp = keyModification.Timestamp
		var t = timeStamp.Seconds*1000 + int64(timeStamp.Nanos/1000000)
		var translated = KeyModification{
			keyModification.TxId,
			keyModification.Value,
			TimeLong(t),
			keyModification.IsDelete}
		if filter(translated) {
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

func ParseStates(iterator shim.StateQueryIteratorInterface) []StateKV {
	defer iterator.Close()
	var kvs []StateKV
	for iterator.HasNext() {
		kv, err := iterator.Next()
		PanicError(err)
		kvs = append(kvs, StateKV{kv.Namespace, kv.Key, string(kv.Value)})
	}
	return kvs
}

type Modifier func(interface{})
type Filter func(interface{}) bool

type Args struct {
	buff [][]byte
}

func ArgsBuilder(fcn string) (Args) {
	return Args{[][]byte{[]byte(fcn)}}
}

func (t *Args) AppendBytes(bytes []byte) {
	t.buff = append(t.buff, bytes)
}
func (t *Args) AppendArg(str string) {
	t.buff = append(t.buff, []byte(str))
}
func (t Args) Get() [][]byte {
	return t.buff
}

//a readable structure of peer.response
type PeerResponse struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `json:"status,omitempty"`
	// A message associated with the response code.
	Message string `json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload string `json:"payload,omitempty"`
}
