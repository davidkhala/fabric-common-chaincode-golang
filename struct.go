package golang

import (
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type StringList struct {
	Strings []string
}

func (s *StringList) construct() {
	if s.Strings == nil {
		s.Strings = []string{}
	}
}
func (s *StringList) Has(value string) bool {
	s.construct();
	for _, e := range s.Strings {
		if e == value {
			return true
		}
	}
	return false
}

func (s *StringList) String() string {
	return strings.Join(s.Strings, ",")
}
func (s *StringList) Put(value string) {
	s.construct();
	s.Strings = append(s.Strings, value)
}

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
	States []KVJson
}

type KVJson struct {
	Namespace string
	Key       string
	Value     string
}

func (state *States) ParseStates(iterator shim.StateQueryIteratorInterface) {
	defer iterator.Close()
	var kvs []KVJson
	for iterator.HasNext() {
		kv, err := iterator.Next()
		PanicError(err)
		kvs = append(kvs, KVJson{kv.Namespace, kv.Key, string(kv.Value)})
	}
	state.States = kvs
}

type Modifier func(interface{})
type Filter func(interface{}) bool
