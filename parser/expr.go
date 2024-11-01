package parser

import "github.com/ethan-prime/graphite/tokens"

// define base expr for all other exprs
type Expr interface {}

type ExprNode struct {
	Expr Expr
}

func (parser *Parser) ParseExpression() *ExprNode {
	return &ExprNode{}
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