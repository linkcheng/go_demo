package mydata

import (
	"go/ast"
	"go/parser"
	"go/token"
)


func Test() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

	expr, err := parser.ParseExpr("a > 100 && b < 1000")
	ast.Print(fset, expr)
}