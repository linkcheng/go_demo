package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i:=0; i<10; i++ {
		i := i
		go func() {
			fmt.Println("a: ", i)
		}()
	}
	// var ch = make(chan int)
	// <- ch
	// 先输出 a: 9, 因为最后创建的 goroutine 在 runnext 中，优先级最高；
	// runnext > local run queue > global run queue
	time.Sleep(time.Minute)
}
