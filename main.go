package main

import (
	"fmt"
	"github.com/ethan-prime/graphite/lexer"
	"github.com/ethan-prime/graphite/parser"
	"github.com/ethan-prime/graphite/codegen"	
)

func main() {
	fmt.Println("Welcome to Graphite...")
	fmt.Printf("Filename> ")

	// get filename input from user
	var filename string
	fmt.Scan(&filename)

	// load it into the lexer
	l := lexer.Lexer{Input: nil, LineNumber: 1, Index: 0, ShowDebug: true}
	l.LoadInput(filename)

	p := parser.Parser{Tokens: l.Tokenize(), Index: 0, ShowDebug: true}
	program_node := p.ParseProgram()

	codegen.ProgramCodeGen(*program_node)
}