package mydata

import "fmt"

// 二叉搜索树
type BstNode struct {
	Data int
	Left *BstNode
	Right *BstNode
}

func NewBstNode(data int) *BstNode {
	return &BstNode{Data: data}
}

func (b *BstNode)Insert(data int) {
    if b.Data > data {
        if b.Left != nil {
            b.Left.Insert(data)
        } else {
            b.Left = NewBstNode(data)
        }
    } else {
        if b.Right != nil {
            b.Right.Insert(data)
        } else {
            b.Right = NewBstNode(data)
        }
    }
}

func (b *BstNode) Search(data int) *BstNode {
    if b.Data > data {
        if b.Left != nil {
            return b.Left.Search(data)
        } 
    } else if b.Data < data {
        if b.Right != nil {
            return b.Right.Search(data)
        } 
    } else {
        return b
    }
    return nil
}

func (b *BstNode) Delete(data int) *BstNode {
    if data < b.Data {
		b.Left = b.Left.Delete(data)
	} else if data > b.Data {
		b.Right = b.Right.Delete(data)
	} else {
		// 被删除的就是当前节点
		if b.Left != nil && b.Right != nil {
			// 被删除的节点有两个子节点
			minNode := b.Right.Min()
		    b.Data = minNode.Data
			b.Right = b.Right.Delete(minNode.Data)
		} else {
			// 被删除的节点有0个或者1个子节点
			if b.Left == nil {
				b = b.Right
			} else if b.Right == nil {
				b = b.Left
			}
		}
	}
    return b
}

func (b *BstNode) Min() *BstNode {
    if b == nil {
        return nil
    }
    cur := b
    for cur.Left != nil {
        cur = cur.Left
    }
    return cur
}

type BST struct {
	Root *BstNode
	Len int
}

func NewBst() *BST {
    return &BST{}
}

func (t *BST) Insert(data int) {
	if t.Len == 0 {
        t.Root = NewBstNode(data)
    } else {
        t.Root.Insert(data)
    }
    t.Len ++
}

func (t *BST) Search(data int) *BstNode {
    if t.Len == 0 {
        return nil
    } else {
        return t.Root.Search(data)
    }
}

func (t *BST) Delete(data int) *BstNode {
    node := t.Search(data)
    if node == nil {
        return nil
    }

	t.Root = t.Root.Delete(data)
    t.Len --
    return node
}

// 按层遍历
func (t *BST) LevelOrder() {
	if t == nil {
		return 
	}

	var res []*BstNode
	res = append(res, t.Root)

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
