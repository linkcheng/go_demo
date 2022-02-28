package mydata

import (
	"fmt"
	"strings"
)

// 单向链表
type LNode struct {
	Data interface{}
	Next *LNode
}

type TList struct {
	Head *LNode  // 头，不参与列表整体
	Len int
}


func CreateNode(val interface{}) *LNode {
	return &LNode{val, nil}
}

func CreateList() *TList {
	node := CreateNode(nil)
	return &TList{node, 0}
}

// 头部添加
func (t *TList) Add(data interface{}) {
	defer func() { t.Len ++ }()

	node := CreateNode(data)
	node.Next = t.Head.Next
	t.Head.Next = node
}

// 尾部追加
func (t *TList) Append(data interface{}) {
	defer func() { t.Len ++ }()

	node := CreateNode(data)
	prev := t.Head
	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = node
}

// 插入
func (t *TList) Insert(index int, data interface{}) {
	defer func() { t.Len ++ }()

	if index > t.Len {
		t.Append(data)
		return
	}
	prev := t.Head
	node := CreateNode(data)

	for i:=0; i<index; i++ {
		prev = prev.Next
	}

	node.Next = prev.Next
	prev.Next = node
}

// 删除
func (t *TList) Delete(index int) interface{} {
	if index > t.Len {
		return nil
	}

	defer func() { t.Len -- }()

	prev := t.Head
	for i:=0; i<index; i++ {
		prev = prev.Next
	}

	hit := prev.Next
	prev.Next = hit.Next
	hit.Next = nil

	return hit.Data
}

// 遍历
func (t *TList) Scan() {
	i := 1
	cur := t.Head.Next
	for cur != nil {
		fmt.Printf("第%d个节点是%d\n", i, cur.Data)
		cur = cur.Next
		i++
	}
}

// 遍历
func (t *TList) ToString() string {
	res := []string{}
	cur := t.Head.Next
	for cur != nil {
		res = append(res, fmt.Sprintf("%+v", cur.Data))
		cur = cur.Next
	}
	return strings.Join(res, ",")
}

// 反转
func (t *TList) Reverse() {
	var prev *LNode
	cur := t.Head.Next

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	t.Head.Next = prev
}

// 是否有环
func (t *TList) HasCycle() bool {
	first := t.Head.Next
	if first == nil {
        return false
    }
	
    slow, fast := first, first.Next
    for fast != slow {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
} 
