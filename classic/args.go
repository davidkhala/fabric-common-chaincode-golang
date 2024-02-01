package classic

type Args struct {
	buff [][]byte
}

func ArgsBuilder(fcn string) Args {
	return Args{[][]byte{[]byte(fcn)}}
}

func (t *Args) AppendBytes(bytes []byte) *Args {
	t.buff = append(t.buff, bytes)
	return t
}
func (t *Args) AppendArg(str string) *Args {
	t.buff = append(t.buff, []byte(str))
	return t
}
func (t Args) Get() [][]byte {
	return t.buff
}
