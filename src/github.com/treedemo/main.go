package main

import (
	"fmt"

	"github.com/treedemo/tree"
)

func add(val int) (res int) {
	res = val + 1
	return
}

/*
模块使用步骤
1. go mod init <module name>
2. go install .
3. go run main.go
*/

func main() {
	root := tree.Node{Value: 1}

	root.Left = tree.CreateNode(2)
	root.Right = tree.CreateNode(3)

	root.Left.Left = tree.CreateNode(4)
	root.Left.Right = tree.CreateNode(5)

	root.Traversal()
	fmt.Println()

	fmt.Println(add(5))
}
