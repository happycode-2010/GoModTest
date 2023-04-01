package GoModTest

import "testing"

func TestPrintSome(t *testing.T) {
	in := "hello"
	out := "Hello: hello"

	ret := PrintSome(in)
	if ret != out {
		t.Errorf("TestPrintSome error, in: %s, ret: %s, exp: %s", in, ret, out)
	}
}
