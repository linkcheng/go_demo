package main

import (
	"fmt"
)

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Printf("i=%d, v=%d; ", i, v)
	}
	fmt.Println()
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
}
