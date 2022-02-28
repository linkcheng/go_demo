package mydata

import (
	"fmt"
)

type TreeNode struct {
	Data interface{}
	Left *TreeNode
	Right *TreeNode
}

func CreateTreeNode(data interface{}) *TreeNode {
	return &TreeNode{Data: data}
}

// 递归实现
// 先序遍历，先读根，再读left，然后right
func (t *TreeNode) PreOrder() {
	if t == nil {
		return 
	}
	fmt.Printf("%v ", t.Data)
	t.Left.PreOrder()
	t.Right.PreOrder()
}

// 中序遍历，先读left，再读根，然后right
func (t *TreeNode) InOrder() {
	if t == nil {
		return 
	}
	t.Left.InOrder()
	fmt.Printf("%v ", t.Data)
	t.Right.InOrder()
}


// 后序遍历，先读left，再读right，最后读根
func (t *TreeNode) PostOrder() {
	if t == nil {
		return 
	}
	t.Left.PostOrder()
	t.Right.PostOrder()
	fmt.Printf("%v ", t.Data)
}

// 非递归实现

// 按层遍历
func (t *TreeNode) LevelOrder() {
	if t == nil {
		return 
	}

	var res []*TreeNode
	res = append(res, t)

	for len(res) != 0 {
		root := res[0]
		res = res[1:]
		fmt.Printf("%v ", root.Data)
		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
	}
}

// 先序遍历，先读根，再读left，然后right
func (t *TreeNode) PreOrder2() {
	if t == nil {
		return 
	}
	res := []*TreeNode{t}
	length := len(res)
	cur := t
	for length > 0  {
		cur = res[length-1]
		res = res[:length-1]

		fmt.Printf("%v ", cur.Data)
		if cur.Right != nil {
			res = append(res, cur.Right)
		}
		if cur.Left != nil {
			res = append(res, cur.Left)
		}
		length = len(res)
	}
}

// 中序遍历，先读left，再读根，然后right
func (t *TreeNode) InOrder2() {
	if t == nil {
		return 
	}
	res := []*TreeNode{}
	root := t
	for root != nil || len(res) != 0 {
		if root != nil {
			res = append(res, root)
			root = root.Left
		} else {
			length := len(res)
			root = res[length-1]
			res = res[:length-1]

			fmt.Printf("%v ", root.Data)
			root = root.Right
		}
	}
}


// 后序遍历，先读left，再读right，最后读根
func (t *TreeNode) PostOrder2() {
	if t == nil {
		return 
	}
	helper := []*TreeNode{}
	res := []*TreeNode{t}
	length := len(res)
	root := t
	for length > 0 {
		root = res[length-1]
		res = res[:length-1]
		helper = append(helper, root)

		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
		length = len(res)
	}

	length = len(helper)
	for i := range helper {
		fmt.Printf("%v ", helper[length-1-i].Data)
	}
}
