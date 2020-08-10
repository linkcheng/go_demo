package main

import (
	"fmt"
	"reflect"

	"github.com/stackdemo/stack"
)

// 判断 Stack 是否实现 IStack
var _ stack.IStack = new(stack.Stack)

func main() {
	data := stack.New(3)

	fmt.Println("Push(1)", data.Push(1))
	fmt.Println("Cap=", data.Cap())
	fmt.Println("Peek=", data.Peek())
	fmt.Println("Len=", data.Len())

	fmt.Println("Push(2)", data.Push(2))
	fmt.Println("Peek=", data.Peek())
	fmt.Println("Len=", data.Len())

	fmt.Println("Pop=", data.Pop())
	fmt.Println("Peek=", data.Peek())
	fmt.Println("Len=", data.Len())

	fmt.Println("Push(2)", data.Push(2))
	fmt.Println("Push(3)", data.Push(3))
	fmt.Println("Push(4)", data.Push(4))

	t := reflect.TypeOf(data)
	fmt.Println("type=", t)
	fmt.Println("kind=", t.Kind())

}
