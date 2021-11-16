package main

import (
	"fmt"
)

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	id int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Bey says, hello my name is %s\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, hello my name is %s\n", j.name)
}


func main3() {
	var ben = &Ben{name: "Ben"}
	var jerry = &Jerry{name: "Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()
	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}
	go loop0()

	for {
		maker.Hello()
	}
}
