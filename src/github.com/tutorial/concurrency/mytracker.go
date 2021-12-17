package main

import (
	"context"
	"fmt"
	"time"
)

/*
1. 把并发扔给调用者控制
2. 搞清楚这个 gouroutine 什么时候退出，管控其生命周期
3. 可以主动控制 channel 什么时候退出，channel关闭close/stop标志 & context超时两种方式
*/

func main5() {
	tr := NewTracker()
	// 并行的行为丢给调用者
	go tr.Run()
	
	ctx := context.Background()
	_ = tr.Event(ctx, "test1")
	_ = tr.Event(ctx, "test2")
	_ = tr.Event(ctx, "test3")

	// 控制退出
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

type Tracker struct {
	ch chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

// 用 channel 处理消息，通过 context 管控生命周期
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case <-ctx.Done():  // 超时控制
		return ctx.Err()
	case t.ch <- data:  // 数据写入 channel
		return nil
	}
}

// 消费 channel 
func (t *Tracker) Run()  {
	for data := range t.ch {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	// 控制退出时机
	t.stop <- struct{}{}
}

// goroutine 退出时机
func (t *Tracker) Shutdown(ctx context.Context) {
	// 主动控制 Run 方法退出
	close(t.ch)

	select {
	case <-t.stop:  // 任务完成，退出
	case <-ctx.Done():  // 超时处理
	}
}
