package encode

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	var post Post
	err := encode("test.json", post)
	if err != nil {
		t.Error(err)
	}
}
func TestDecode(t *testing.T) {
	post, err := decode("test.json")
	if err != nil {
		t.Fail()
	}
	fmt.Println(post)
}

func TestAdvanceMethod(t *testing.T) {
	//该方法还没有实现是要跳过的
	t.Skip("not implemented")
}
func TestMethodWithLongTime(t *testing.T) {
	if testing.Short() {
		t.Skip("skip because long")
	}
	methodWithLongTime()
}
