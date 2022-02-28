package mydata

import (
	"fmt"
	"testing"
)

/*
      1
    /   \
  2       3
 / \    /  \
4   5  6    7
      / \
     8   9

前序输出: 1 2 4 5 3 6 8 9 7
中序输出: 4 2 5 1 8 6 9 3 7
后序输出: 4 5 2 8 9 6 7 3 1
层序输出: 1 2 3 4 5 6 7 8 9
*/
func Test_TreeNode(t *testing.T) {
	root := CreateTreeNode(1)
	root.Left = CreateTreeNode(2)
	root.Left.Left = CreateTreeNode(4)
	root.Left.Right = CreateTreeNode(5)

	root.Right = CreateTreeNode(3)
	root.Right.Left = CreateTreeNode(6)
	root.Right.Left.Left = CreateTreeNode(8)
	root.Right.Left.Right = CreateTreeNode(9)
	root.Right.Right = CreateTreeNode(7)

	root.PreOrder()
	fmt.Println()

	root.InOrder()
	fmt.Println()

	root.PostOrder()
	fmt.Println()

	fmt.Println("===================")

	root.PreOrder2()
	fmt.Println()

	root.InOrder2()
	fmt.Println()

	root.PostOrder2()
	fmt.Println()

	root.LevelOrder()
	fmt.Println()
}
