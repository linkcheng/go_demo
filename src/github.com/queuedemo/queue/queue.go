package queue

import "container/list"

/*
Queue 队列
*/
type Queue struct {
	data *list.List
	cap  int
}

/*
Element 队列元素
*/
type Element interface{}

/*
IQueue 队列接口
*/
type IQueue interface {
	Enqueue(e Element) bool
	Dequeue() Element
	Peek() Element
	IsEmpty() bool
	IsFull() bool
	Len() int
	Cap() int
}

/*
New 创建队列
*/
func New(cap int) Queue {
	queue := Queue{&list.List{}, cap}
	return queue
}

/*
Enqueue 入队
*/
func (q *Queue) Enqueue(e Element) bool {
	if q.data.Len() >= q.cap {
		return false
	}
	q.data.PushBack(e)
	return true
}

/*
Dequeue 出队
*/
func (q *Queue) Dequeue() Element {
	if q.data.Len() <= 0 {
		return nil
	}
	return q.data.Remove(q.data.Front())
}

/*
Peek 查看队首元素
*/
func (q *Queue) Peek() Element {
	if q.cap <= 0 {
		return nil
	}
	return q.data.Front().Value
}

/*
IsEmpty 查看是否是空队列
*/
func (q *Queue) IsEmpty() bool {
	return q.cap == 0
}

/*
IsFull 查看是否是满队列
*/
func (q *Queue) IsFull() bool {
	return q.cap == q.data.Len()
}

/*
Len 查看队列中元素数量
*/
func (q *Queue) Len() int {
	return q.data.Len()
}

/*
Cap 查看栈的容量
*/
func (q *Queue) Cap() int {
	return q.cap
}
