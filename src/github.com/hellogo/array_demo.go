package main

import (
	"container/list"
	"fmt"
	"math"
)

type myFloat float64

var container = []string{"zero", "one", "two"}

func (num myFloat) abs() float64 {
	return math.Abs(float64(num))
}

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Printf("i=%d, v=%d; ", i, v)
	}
	fmt.Println()
}

/*
calculate...
计算器 + — * /
*/
func calculate(input string) int {
	stack := list.New()
	numStr := ""

	for _, s := range input {
		fmt.Printf("val=%v, type=%T, ", s, s)
		switch s {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			numStr := numStr + string(s)
		case '+', '-', '*', '/':
			op := stack.Back()
			if op != nil && (op == "*" || op == "/") {
				op = stack.Remove(stack.Back())
				prev := stack.Remove(stack.Back())
			}

			num1 := int(prev) / int(num)

			stack.PushBack(numStr)
			stack.PushBack(string(s))
		}
	}

	return 0
}

func updateArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 10
	}
}

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	printArray(a1[:])
	updateArray(a1[:])
	printArray(a1[:])

	a2 := a1[1:5]
	printArray(a2)
	fmt.Printf("len(a2)=%d, cap(a2)=%d\n", len(a2), cap(a2))

	a3 := a2[3:6]
	printArray(a3)
	fmt.Printf("len(a3)=%d, cap(a3)=%d\n", len(a3), cap(a3))

	value, ok := interface{}(container).([]string)
	fmt.Println(value, ok)

	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	value1, ok1 := interface{}(container).(map[int]string)
	fmt.Println(value1, ok1)

	value2, ok2 := interface{}(container).([]string)
	fmt.Println(value2, ok2)

	// mylist := list.List()
	fmt.Printf("T(new([]int)=%T\n", new([]int))
	fmt.Printf("T(make([]int, 10, 20)=%T\n", make([]int, 10, 20))

	calculate("123123")
}
