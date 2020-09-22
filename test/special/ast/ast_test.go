package main_test

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"testing"
)

func TestAst_Parse(t *testing.T) {
	expr, _ := parser.ParseExpr("a * -1")
	fmt.Printf("%#v\n", expr)
}

func TestHello(t *testing.T) {
	fmt.Printf("Hello")

	var tk token.Token
	_ = tk

	var f token.FileSet
	_ = &f

	var sc scanner.Scanner
	_ = sc
}

type IMy interface {
	Val() string
}

type hello struct{}
