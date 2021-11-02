package main

import "fmt"

// Person ...
type Person struct {
	Name   string
	Sexual string
	Age    int
}

// Print ...
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func test2() {
	var p = Person{
		Name:   "Long Zheng",
		Sexual: "Male",
		Age:    18,
	}
	p.Print()
}

// PrintPerson ...
func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func test1() {
	var p = Person{
		Name:   "Long Zheng",
		Sexual: "Male",
		Age:    18,
	}
	PrintPerson(&p)
}

// Shape ...
type Shape interface {
	Sides() int
	Area() int
}

// Square ...
type Square struct {
	len int
}

// Sides ...
func (s *Square) Sides() int {
	return 4
}

// Area ...
func (s *Square) Area() int {
	return s.len * s.len
}

func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())

	var _ Shape = (*Square)(nil)
}
