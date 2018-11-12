package golang

import (
	"fmt"
	"testing"
)

func TestArgsBuilder_Constructor(t *testing.T) {
	var args = ArgsBuilder("abc")
	fmt.Println("args", args)
	args.AppendArg("cde")
	args.AppendBytes(nil)
	var result = args.Get()
	fmt.Println("args2", result)
	fmt.Println(`args[2]==""`, string(result[2]) == "")
}
