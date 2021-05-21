package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	var funcname = "parserfile"
	a, err := parserNew(funcname)

	if err != nil {
		panic(err)
	}

	fmt.Println("Количество вызовов асинхронных функций:", a, err)
}

func parserNew(funcName string) (int32, error) {

	var a int32

	fset := token.NewFileSet()
	src, err := os.Open(funcName + ".go")
	_ = err
	f, err := parser.ParseFile(fset, funcName, src, 0)

	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.GoStmt:
			a++
		}
		return true
	})

	return a, err
}
