package empty_interface_test

import (
	"fmt"
	"testing"
)

func Dosomething(p interface{}) {
	cc := p.(int)
	fmt.Println("type", cc)

	// if i, ok := p.(int); ok {
	// fmt.Println("Integer", i)
	// return
	// }

	// if s, ok := p.(string); ok {
	// 	fmt.Println("String", s)
	// 	return
	// }

	// fmt.Println("Unknow Type")

	// switch v := p.(type) {
	// case int:
	// 	fmt.Println("Integer", v)
	// case string:
	// 	fmt.Println("String", v)
	// default:
	// 	fmt.Println("Unknow Type")

	// }

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	Dosomething(10)
	Dosomething("10")
}
