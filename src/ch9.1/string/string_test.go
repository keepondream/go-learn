package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xBF\xA5\xFF"
	t.Log(s)
	t.Log(len(s))
	// 字符串长度统计的是 byte数
	// s[1] = "3" // string是不可变的byte slice
	s = "中国"
	t.Log(s)
	t.Log(len(s)) // byte数

	c := []rune(s)
	t.Log(c)
	t.Log(len(c))
	t.Logf("中国 Unicode %x %x %U ", c[1], c[0], c[0])
	t.Logf("中国 utf8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, value := range s {
		t.Logf("%c", 22269)
		t.Log(value)
		t.Logf("%[1]c %[1]x", value)
	}

	t.Log(len([]rune(s)))
}

// golang 中的 字符串 string 用 byte 表示
// byte 等同于 unit8 ,通常用来处理 ASCII字符
// rune 等同于 int32 ,通常用来处理 unicode 或 utf-8字符
