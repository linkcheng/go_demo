// Practice
// 1、嵌套 interface
// 2、嵌套 struct
// 注意指针的使用
package main

import (
	"log"
)

type InfParent interface {
	runParent()
}

// inherit InfParent interface
type InfChild interface {
	InfParent
	runChild()
}

// parent struct
type StParent struct {
	name string
}

// inherit parent struct
type StChild struct {
	InfChild
	*StParent
	age int
}

// implementation InfChild interface
type ZR int

func (z *ZR) runChild() {
	log.Println("runChild")
}
func (z *ZR) runParent() {
	log.Println("runParent")
}

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

// main
func main() {
	var zr ZR
	zr = 1

	var stp StParent
	stp = StParent{name: "StParent.name"}

	st := &StChild{
		InfChild: &zr,
		// wrong
		// ZR does not implement InfChild (runChild method has pointer receiver)
		//InfChild: zr,
		age:      2,
		StParent: &stp,
	}
	log.Printf("%#v", st)

	// 直接使用继承的属性
	log.Println(st.name)

	// 访问方法一
	log.Println("--> Call One")
	st.runChild()
	st.runParent()

	// 访问方法二
	log.Println("--> Call Two")
	st.InfChild.runChild()
	st.InfChild.runParent()
}

// output
// embed.go:62: &main.StChild{InfChild:(*main.ZR)(0xc000014088), StParent:(*main.StParent)(0xc000010200), age:2}
// embed.go:65: StParent.name
// embed.go:68: --> Call One
// embed.go:35: runChild
// embed.go:38: runParent
// embed.go:73: --> Call Two
// embed.go:35: runChild
// embed.go:38: runParent
