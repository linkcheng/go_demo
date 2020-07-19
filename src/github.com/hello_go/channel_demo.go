package main

import (
	"fmt"
	"reflect"
	)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB

)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}

func finonacci_test()  {
	c := make(chan int)
	quit := make(chan int)

	go func () {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func main() {
	fmt.Println("hello world")

	for pos, char := range "aΦx" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
	//fmt.Println(B)
	//fmt.Println(KB)
	//fmt.Println(MB)
	//fmt.Println(GB)

	//a := [...]int{1,2,3,4,5}
	//sa := a[:]
	//sc := a[:3]
	//
	//fmt.Println(a)
	//fmt.Println(sa)
	//fmt.Println(sc)
	//
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(sa))
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p\n", &sa)
	//fmt.Printf("%p\n", &sc)
	//
	//sa[0] = 10
	//fmt.Println(a)
	//fmt.Println(sa)
	//fmt.Println(sc)

	a := [...]int{1,2,3,4,5}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(len(a), cap(a))

	nums := []int{2, 3, 4}
	fmt.Println(reflect.TypeOf(nums))
	fmt.Println(len(nums), cap(nums))
	
	m := make([]int, 4, 8)
	fmt.Println(m)
}
