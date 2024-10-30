package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Vulcan...")
	fmt.Printf("Filename> ")

	// get filename input from user
	var filename string
	fmt.Scan(&filename)

	// load it into the lexer
	lexer := Lexer{input: nil, line_number: 0, index: 0, show_debug: true}
	lexer.LoadInput(filename)
	lexer.Tokenize()
}