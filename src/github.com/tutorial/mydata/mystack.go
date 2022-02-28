package mydata

import (
	"sync"
)

type SNode struct {
	Data interface{}
	Prev *SNode
}

type Stack struct {
	Top *SNode
	Len int
	Lock   *sync.RWMutex
}

func CreateSNode(value interface{}, prev *SNode) *SNode {
	return &SNode{value, prev}
}

func CreateStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.Len == 0 {
		return nil
	}
	return s.Top.Data
}

// 压栈
func (s *Stack) Pop() interface{} {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.Len == 0 {
		return nil
	}
	defer func(){ s.Len-- }()

	cur := s.Top
	s.Top = cur.Prev

	cur.Prev = nil
	return cur.Data
}

// 压栈
func (s *Stack) Push(value interface{}) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	defer func(){ s.Len++ }()

	node := CreateSNode(value, s.Top)
	// node.Prev = s.Top
	s.Top = node
}

func (s *Stack) Empty() bool {
    return s.Len == 0
}


type MyStack struct {
    Data []int
}

func Constructor() MyStack {
    return MyStack{[]int{}}
}

func (s *MyStack) Push(x int)  {
    s.Data = append(s.Data, x)
}

func (s *MyStack) Pop() int {
    len := len(s.Data)
    val := s.Data[len-1]
    s.Data = s.Data[:len-1]
    return val
}

func (s *MyStack) Top() int {
    len := len(s.Data)
    val := s.Data[len-1]
    return val
}

func (s *MyStack) Empty() bool {
    return len(s.Data) == 0
}


func ChechParenthese(s string) bool {
	if len(s) % 2 == 1 {
        return false
    }

	charMap := map[rune]rune{
		')': '(',
		']': '[', 
		'}': '{',
	} 
	stack := CreateStack()
	for _, r := range s {
		if r == '(' ||  r == '[' ||  r == '{' {
			stack.Push(r)
		} else {
			if stack.Empty() {
				return false
			}
			top := stack.Pop().(rune)
			if top != charMap[r] {
				return false
			}
		}
	}

	return stack.Empty()
}
