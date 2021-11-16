package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 超时控制
type result struct {
	record string
	err error
}

func search(term string) (string, error) {
	// time.Sleep(50*time.Millisecond)
	time.Sleep(200*time.Millisecond)
	return "search term", nil
}


func process() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelFunc()

	ch := make(chan result)

	go func() {
		record, err := search("aaa")
		ch <- result{record, err}
	}()

	select {
	case <-ctx.Done():  // 超时控制
		fmt.Println("Canceled")
		return errors.New("search time out: canceled")
	case result := <-ch:  // 明确推出时机
		if result.err != nil {
			fmt.Println(result.err.Error())
			return result.err
		}
		fmt.Println("Receive:", result.record)
		return nil
	}
}

func main1() {
	process()
}
