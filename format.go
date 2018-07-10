package golang

import (
	"strconv"
	"encoding/json"
	"errors"
)

func ToInt(bytes []byte) int {
	if bytes == nil {
		return 0
	}
	i, err := strconv.Atoi(string(bytes))
	PanicError(err)
	return i
}
func ToBytes(integer interface{}) []byte {
	return []byte(strconv.FormatInt(integer.(int64), 10))
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
func PanicString(err string) {
	if err != "" {
		panic(errors.New(err))
	}
}

/**
	a wrapper to panic Unmarshal(non-pointer v)
 */
func FromJson(jsonString []byte, v interface{}) {
	err := json.Unmarshal(jsonString, v)
	PanicError(err);
}

func ToJson(v interface{}) []byte {
	result, err := json.Marshal(v)
	PanicError(err)
	return result
}
