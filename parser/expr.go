package parser

import (
	"github.com/ethan-prime/graphite/tokens"
)

// define base expr for all other exprs
type Expr interface {}

type ExprNode struct {
	Expr Expr
}

// parses an (arbitrarily long) expression
func (parser *Parser) ParseExpression() *ExprNode {
	LHS := parser.ParsePrimaryExpression()

	// parse RHS, if it exists...
	return parser.ParseBinaryOpRHS(0, LHS)
}

// Parses RHS of a binary expression.
// Operator-Precedence Parsing Algorithm
func (parser *Parser) ParseBinaryOpRHS(current_precedence int, LHS *ExprNode) *ExprNode {
	cur_tok := parser.CurrentToken()

	for { // recursively parse RHS
		precedence := parser.OperatorPrecedence(cur_tok.Value)

		if precedence < current_precedence {
			return LHS
		}

		// so we must have a binary operator
		binop := parser.CurrentToken()
		parser.Advance()

		RHS := parser.ParsePrimaryExpression()

		next_precedence := parser.OperatorPrecedence(parser.CurrentToken().Value)
		if precedence < next_precedence {
			RHS = parser.ParseBinaryOpRHS(precedence + 1, RHS)
		}

		// merge LHS and RHS
		LHS = &ExprNode{
			Expr: &BinaryExprNode{operator: binop.Value, LHS: LHS, RHS: RHS},
		}

	}
}

// parses a parenexpr
// <paren_expr> ::= ( expr )
func ParseParenExpr(parser *Parser) *ExprNode {
	cur_tok := parser.CurrentToken()
	if cur_tok.ID != tokens.OPEN_PAREN {
		parser.ParserError("ParseParenExpr", "(", cur_tok.Repr(), cur_tok.LineNumber)
	}
	parser.Advance()
	
	// get expression
	expr := parser.ParseExpression()

	parser.Advance()

	cur_tok = parser.CurrentToken()
	if cur_tok.ID != tokens.OPEN_PAREN {
		parser.ParserError("ParseParenExpr", ")", cur_tok.Repr(), cur_tok.LineNumber)
	}

	return expr
}