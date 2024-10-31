package parser

import (
	"log"

	"github.com/ethan-prime/graphite/tokens"
)

type PrimaryExpr interface {}

type PrimaryExprNode struct {
	PrimaryExpr PrimaryExpr
}

type DoubleLiteral struct {
	Value string
}

type VariableReference struct {
	Value string
}

/*
type FunctionCall struct {
	FunctionName string
	Arguments    []ExprNode
}
*/

// implements ParsePrimaryExpr for DoubleLiteral
func (parser *Parser) ParsePrimaryExpr() (primary_expr PrimaryExprNode) {
	cur_tok := parser.CurrentToken()
	if cur_tok.ID == tokens.DOUBLE {
		// advance the parser, return the parsed literal
		parser.Advance()
		primary_expr = PrimaryExprNode{DoubleLiteral{cur_tok.Value}}
	} else if cur_tok.ID == tokens.IDENTIFIER {
		parser.Advance()
		primary_expr = PrimaryExprNode{VariableReference{cur_tok.Value}}
	} else {
		log.Fatalf("[graphite PARSER] ParsePrimaryExpr():\n\tExpected: Primary Expression\n\tReceived: %s\n", cur_tok.Repr())
	}
	return
}