package main

import (
	"fmt"
	"reflect"

	"github.com/queuedemo/queue"
)

func main() {
	q := queue.New(3)
	for i := 1; i < 5; i++ {
		fmt.Println("Enqueue", q.Enqueue(i))
	}

	fmt.Println("Len", q.Len())
	fmt.Println("Peek", q.Peek())
	fmt.Println("Dequeue", q.Dequeue())
	fmt.Println("Dequeue", q.Dequeue())
	fmt.Println("Len", q.Len())

	t := reflect.TypeOf(q)
	fmt.Println("type=", t)
	fmt.Println("kind=", t.Kind())
}
