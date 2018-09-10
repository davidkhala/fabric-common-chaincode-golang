package golang

import (
	"testing"
	"fmt"
)

func TestArgsBuilder_Constructor(t *testing.T) {
	var args = ArgsBuilder("abc")
	fmt.Println("args",args)
	args.AppendArg("cde")
	fmt.Println("args2",args)

}