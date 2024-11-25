package string_func

import "testing"
import "strings"
import "strconv"

func TestStringSplit(t *testing.T) {
	s := "A,B,C,D"
	parts := strings.Split(s,",")
	for _,part := range parts {
		t.Log(part)
	}
	news := strings.Join(parts, "-")
	t.Log(news)
}

func TestStringConv(t *testing.T) {
	a := 10
	s := strconv.Itoa(a)
	t.Logf("s = %s", s)
	if b,err := strconv.Atoi("12"); err == nil {
		t.Logf("a+b=%d", a+b)
	}
}