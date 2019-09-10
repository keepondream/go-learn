package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	var (
		d int = 1
		f     = 2
	)
	t.Log(a, b, d, f)

	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}

func TestFibList1(t *testing.T) {
	a, b := 1, 1

	for i := 0; i < 9; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(a, b)
}
