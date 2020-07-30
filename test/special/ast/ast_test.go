package ast_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"
)

func TestAst_Parse(t *testing.T) {
	expr, _ := parser.ParseExpr("a * -1")
	fmt.Printf("%#v\n", expr)
}

func TestHello(t *testing.T) {
	fmt.Printf("Hello")
	token.Token
	token.FileSet

}
