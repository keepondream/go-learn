package empty_interface111

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {

	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

func DoSomethingTwo(p interface{}) {
	switch s := p.(type) {
	case int:
		fmt.Println("switch integer", s)
	case string:
		fmt.Println("switch string", s)
	default:
		fmt.Println("switch default unknow")
	}

}

func TestEmptyInterfaceAssert(t *testing.T) {
	DoSomething(10)
	DoSomething("ccc")

	DoSomethingTwo(111)
	DoSomethingTwo("asdfasdf")
}
