package str

import (
	"testing"
)

func TestSubStr(t *testing.T) {
	str := "hello, world!"
	start := 0
	end := 5
	expect := "hello"
	actual := SubStr(str, start, end)
	if expect != actual {
		t.Errorf("expect %s, but had %s", expect, actual)
	}
}

func BenchmarkSubStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SubStr("hello, world!", 0, 5)
	}
}
