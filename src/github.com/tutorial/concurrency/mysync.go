package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Counter int = 0

func main2() {
	for i:=0; i<2; i++ {
		wg.Add(1)
		go Routine(i)
	}
	wg.Wait()
	fmt.Printf("Final Counter=%d\n", Counter)
}

func Routine(id int) {
	defer wg.Done()
	for i:=0; i<2; i++ {
		v := Counter
		time.Sleep(1*time.Nanosecond)
		v++
		Counter = v
	}
}