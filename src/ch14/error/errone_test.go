package error_test

import (
	"errors"
	"testing"
)

var LessThanTwoError1 = errors.New("n should be not less than 2")

var LargerThenHundredError1 = errors.New("n should be not larger than 100")

func GetFibonaccicc(n int) ([]int, error) {

	if n < 2 {
		return nil, LessThanTwoError1
	}

	if n > 100 {
		return nil, LargerThenHundredError1
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonaccicc(t *testing.T) {
	if v, err := GetFibonaccicc(1); err != nil {
		if err == LessThanTwoError1 {
			t.Log("it is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
