package constant_test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Executable = 1 << iota
	Writable
	Readable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func TestConstantTry1(t *testing.T) {
	// a := 7
	a := 5
	t.Log(Readable, Writable, Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 逻辑运算符 & 按位与运算时双目运算符,功能是参与运算的两位数各对应的二进位相与

	// & 与 两者都为真 则为真
	// | 或 两者有一为真 则为真
	// ^ 异或 两者值不相同 则为真

	t.Log(60 & 13)
}
