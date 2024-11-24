package string_test 

import "testing"

func TestString(t *testing.T) {
	var s string 
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s = "\xE4\xB8\xA5"
	s = "\xE4\xBA\xBB\xFF"
	//s[1] = '3' //cannot assign to s[1] (neither addressable nor a map index expression)
	t.Log(s)
	t.Logf("len(s) = %d", len(s))

	s = "中"
	t.Logf("len(s) = %d", len(s))
	c := []rune(s)
	t.Logf("len(c) = %d", len(c))
	t.Logf("%s unicode %x", s, c)
	t.Logf("%s utf8 %x", s, s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _,ele := range s {
		t.Logf("%[1]c, %[1]x", ele) //%c表示以字符格式输出,%x表示输出unicode
	}
}