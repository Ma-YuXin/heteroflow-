package assemblyslicer

import (
	"testing"
)

func TestCallInstArgs(t *testing.T) {
	s := `  6af129:	e8 72 8f db ff       	callq  4680a0 <runtime.morestack_noctxt.abi0>`
	config := NewConfig()
	res, err := config.extract.callInstArgs(s)
	if err != nil {
		t.Error("CallInstArgs can't running right")
	}
	t.Log(res)
}
