package main

import "fmt"


type BaseT struct {
	Id int64
	Name string
}

func (b *BaseT) String() {
	fmt.Printf("BaseT{Id=%d, Name=%s}", b.Id, b.Name)
}

type S1 struct {
	BaseT
	Age int
}

func main() {
	s1 := S1{BaseT: BaseT{Id: 1, Name: "s1"}, Age: 1}
	s1.String()
}
