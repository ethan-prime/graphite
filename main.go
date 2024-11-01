package main

import (
	"fmt"
	"github.com/ethan-prime/graphite/lexer"
	"github.com/ethan-prime/graphite/parser"
)

func main() {
	fmt.Println("Welcome to Graphite...")
	fmt.Printf("Filename> ")

	// get filename input from user
	var filename string
	fmt.Scan(&filename)

	// load it into the lexer
	lexer := lexer.Lexer{Input: nil, LineNumber: 1, Index: 0, ShowDebug: true}
	lexer.LoadInput(filename)

	parser := parser.Parser{Tokens: lexer.Tokenize(), Index: 0, ShowDebug: true}
	parser.ParseProgram()
}