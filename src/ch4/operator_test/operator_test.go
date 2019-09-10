package operator_test

import "testing"

const (
	Executable = 1 << iota
	Writable
	Readable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == d)
	t.Log(a == b)
}

func TestBitClear(t *testing.T) {
	a := 7 //0111

	a = a &^ Readable
	a = a &^ Executable
	t.Log(a)

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
