package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "a,b,v,c,x"
	arrs := strings.Split(s, ",")
	t.Log(arrs)
	for _, v := range arrs {
		t.Log(v)
	}
	t.Log(strings.Join(arrs, "--)"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string sss " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
