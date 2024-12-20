package parser

import (
	"log"
	"fmt"
	"github.com/ethan-prime/graphite/tokens"
)

type ProgramNode struct {
    Stmts []Stmt
}

type Parser struct {
	Tokens []tokens.Token // list of tokens from lexer
	Index  int            // keep track of which token we're on
	ShowDebug bool
}

// top ::= definition | expression | ';'
func (parser *Parser) ParseProgram() *ProgramNode {
	if parser.Tokens == nil {
		log.Fatalf("Tokens not loaded to Parser.")
	}

    program_node := ProgramNode{}

	for parser.CurrentToken().ID != tokens.EOF {
		program_node.Stmts = append(program_node.Stmts, parser.ParseStatement())
	}

	fmt.Println("Successfully parsed program!")
	return &program_node
}

func (parser *Parser) ParserError(function_name string, expected string, received string, line_number int) {
	log.Fatalf("[ graphite compiler ] parsing error @ %s():\n\tExpected: %s\n\tReceived: %s (line %d)\n", function_name, expected, received, line_number)
}

// advance the Parser to the next token
func (parser *Parser) Advance() {
	parser.Index++
}

// load tokens into Parser
func (parser *Parser) LoadTokens(tokens []tokens.Token) {
	parser.Tokens = tokens
}

// return curent token in Parser
func (parser *Parser) CurrentToken() tokens.Token {
	if parser.Tokens == nil {
		log.Fatal("No tokens loaded into parser...")
	}

	if parser.Index >= len(parser.Tokens) {
		return tokens.Token{ID: tokens.EOF}
	}

	return parser.Tokens[parser.Index]
}

// return next (peek) token in Parser
func (parser *Parser) PeekToken() tokens.Token {
	if parser.Tokens == nil {
		log.Fatal("No tokens loaded into parser...")
	}

	if parser.Index+1 >= len(parser.Tokens) {
		return tokens.Token{ID: tokens.EOF}
	}

	return parser.Tokens[parser.Index+1]
}
