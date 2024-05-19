package assemblyslicer

import (
	"testing"
)

func TestCallInstArgs(t *testing.T) {
	// s := `  6af129:	e8 72 8f db ff       	callq  4680a0 <runtime.morestack_noctxt.abi0>`
	s := ` 80487b8:	e8 25 02 00 00       	call   80489e2 <__get_pc_thunk_bx>`
	extract = IntelExtract{}
	res, err := extract.callInstArgs(s)
	if err != nil {
		t.Error("CallInstArgs can't running right")
	}
	t.Log(res)
}

func TestFunctionName(t *testing.T) {
	s:="0000000000003020 <.plt>:"
	// s := "Disassembly of section .plt:"
	extract = IntelExtract{}
	res, err := extract.functionName(s)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

}
func TestVerb(t *testing.T) {
	// s:=" 80487d0:	ff 35 04 b0 04 08    	pushl  0x804b004"
	s := "	..."
	extract = IntelExtract{}
	res, err := extract.verb(s)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
