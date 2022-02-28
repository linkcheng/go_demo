package mydata

import (
	"fmt"
	"testing"
)

/*
	6
   / \
  2   8
 / \
1   5
   /
  3
   \
    4
*/
func Test_BST(t *testing.T) {
	bst := NewBst()
	bst.Insert(6)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(4)

	bst.LevelOrder()
	fmt.Println("\n=============")

	n := bst.Search(2)
	fmt.Printf("search data = %d\n", n.Data)

	n = bst.Delete(2)
	fmt.Printf("delete data = %d\n", n.Data)
	bst.LevelOrder()
	fmt.Println("\n=============")
}
