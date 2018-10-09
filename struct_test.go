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
	fmt.Println("args2", args)
}
