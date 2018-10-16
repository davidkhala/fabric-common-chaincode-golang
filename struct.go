package golang

import (
	. "github.com/davidkhala/goutils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type KeyModification struct {
	TxId      string
	Value     []byte
	Timestamp int64
	IsDelete  bool
}
type History struct {
	Modifications []KeyModification
}

func (history *History) ParseHistory(iterator shim.HistoryQueryIteratorInterface, filter Filter) {
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
			t,
			keyModification.IsDelete}
		if filter(translated) {
			result = append(result, translated)
		}

	}
	history.Modifications = result
}

type States struct {
	States []StateKV
}

type StateKV struct {
	Namespace string
	Key       string
	Value     string
}

func ParseStates(iterator shim.StateQueryIteratorInterface) States {
	defer iterator.Close()
	var kvs []StateKV
	for iterator.HasNext() {
		kv, err := iterator.Next()
		PanicError(err)
		kvs = append(kvs, StateKV{kv.Namespace, kv.Key, string(kv.Value)})
	}
	return States{kvs}
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
