package main

import (
	"fmt"
	"sync"
	"time"
)

// func hello(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Println(request.Method, request.URL)

// 	go func() {
// 		for range time.Tick(time.Second) {
// 			fmt.Println("Current request is in progress")
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// 	writer.Write([]byte("Hi"))
// }

func test1() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("child goroutine bootstrap start")
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker .")
			case <-quit:
				fmt.Println("work well .")
				ticker.Stop()
				return
			}
		}
		fmt.Println("child goroutine bootstrap end")
	}()
	time.Sleep(10 * time.Second)
	quit <- 1
	wg.Wait()
}

func test2() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

// 等待所有 goroutine 执行完毕
// 使用传址方式为 WaitGroup 变量传参
// 使用 channel 关闭 goroutine

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, ch, done, &wg) // wg 传指针，doIt() 内部会改变 wg 的值
	}

	for i := 0; i < workerCount; i++ { // 向 ch 中发送数据，关闭 goroutine
		ch <- i
	}

	close(done)
	wg.Wait()
	close(ch)
	fmt.Println("all done!")
}

func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	for {
		select {
		case m := <-ch:
			fmt.Printf("[%v] m => %v\n", workerID, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerID)
			return
		}
	}
}
