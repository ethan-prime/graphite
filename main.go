package main

import (
	"fmt"
	"github.com/ethan-prime/graphite/lexer"
)

func main() {
	fmt.Println("Welcome to Vulcan...")
	fmt.Printf("Filename> ")

	// get filename input from user
	var filename string
	fmt.Scan(&filename)

	// load it into the lexer
	lexer := lexer.Lexer{Input: nil, LineNumber: 1, Index: 0, ShowDebug: true}
	lexer.LoadInput(filename)
	lexer.Tokenize()
}