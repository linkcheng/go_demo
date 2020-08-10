package stack

import (
	"container/list"
)

/*
Element ...
*/
type Element interface{}

/*
Stack ...
*/
type Stack struct {
	data *list.List
	cap  int
}

/*
IStack ...
*/
type IStack interface {
	Push(e Element) bool
	Pop() Element
	Peek() Element
	IsEmpty() bool
	IsFull() bool
	Len() int
	Cap() int
}

/*
New 创建
*/
func New(cap int) Stack {
	stack := Stack{&list.List{}, cap}
	return stack
}

/*
Push 数据入栈
*/
func (stack *Stack) Push(e Element) bool {
	if stack.data.Len() >= stack.cap {
		return false
	}
	stack.data.PushBack(e)
	return true
}

/*
Pop 弹出栈顶元素
*/
func (stack *Stack) Pop() Element {
	if stack.cap <= 0 {
		return nil
	}
	return stack.data.Remove(stack.data.Back())
}

/*
Peek 查看栈顶元素
*/
func (stack *Stack) Peek() Element {
	if stack.cap <= 0 {
		return nil
	}
	return stack.data.Back().Value
}

/*
IsEmpty 查看是否是空栈
*/
func (stack *Stack) IsEmpty() bool {
	return stack.cap == 0
}

/*
IsFull 查看是否是满栈
*/
func (stack *Stack) IsFull() bool {
	return stack.cap == stack.data.Len()
}

/*
Len 查看栈中元素数量
*/
func (stack *Stack) Len() int {
	return stack.data.Len()
}

/*
Cap 查看栈的容量
*/
func (stack *Stack) Cap() int {
	return stack.cap
}
