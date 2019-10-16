package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProGrammer struct {
}

func (g *GoProGrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())

}

func TestPolymorphismcc(t *testing.T) {
	goProg := &GoProGrammer{}
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)j
	WriteFirstProgram(javaProg)
}
