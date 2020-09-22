package main

import (
	"fmt"
	"go/ast"
	"go/scanner"
	"go/token"
)

func main() {
	// runToken()
	runAST()
}

func runAST() {
	var lit9527 = &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)
}

func runToken() {

	var src = []byte(`println("你好，世界")
// a is tmp variable
var a bool
`)

	var fset = token.NewFileSet()
	var file = fset.AddFile("hello.go", fset.Base(), len(src))
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
