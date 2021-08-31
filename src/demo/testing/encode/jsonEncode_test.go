package encode

import "testing"

func TestEncode(t *testing.T) {
	if encode() == false {
		t.Error()
	}
}
func TestDecode(t *testing.T) {
	if decode(1000) == false {
		t.Fail()
	}
}
