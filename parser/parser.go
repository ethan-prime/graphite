package parser

import (
	"log"
	"github.com/ethan-prime/graphite/tokens"
)

type Parser struct {
	tokens []tokens.Token // list of tokens from lexer
	index int // keep track of which token we're on
	show_debug bool
}

// advance the Parser to the next token
func (parser *Parser) Advance() {
	parser.index++
}

// load tokens into Parser
func (parser *Parser) LoadTokens(tokens []tokens.Token) {
	parser.tokens = tokens
}

// return curent token in Parser
func (parser *Parser) CurrentToken() tokens.Token {
	if parser.tokens == nil {
		log.Fatal("No tokens loaded into parser...")
	}

	if (parser.index >= len(parser.tokens)) {
		return tokens.Token{ID: tokens.EOF}
	}

	return parser.tokens[parser.index]
}

// return next (peek) token in Parser
func (parser *Parser) PeekToken() tokens.Token {
	if parser.tokens == nil {
		log.Fatal("No tokens loaded into parser...")
	}

	if (parser.index + 1 >= len(parser.tokens)) {
		return tokens.Token{ID: tokens.EOF}
	}

	return parser.tokens[parser.index + 1]
}