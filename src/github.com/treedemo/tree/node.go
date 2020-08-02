package tree

import "fmt"

/*
Node ...
树节点
*/
type Node struct {
	Value       int
	Left, Right *Node
}

/*
Print ...
打印节点信息
*/
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

/*
CreateNode ...
创建节点
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

/*
SetValue ...
设置值
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}

/*
Traversal ...
遍历
*/
func (node *Node) Traversal() {
	if node == nil {
		return
	}

	node.Print()
	node.Left.Traversal()
	node.Right.Traversal()
}
