package parser

import (
	"fmt"

	"github.com/ethan-prime/graphite/tokens"
)

// define base expr for all other exprs
type Expr interface{}

type ExprNode struct {
	Expr Expr
}

// <top_level_expr> ::= <expr>
func (parser *Parser) ParseTopLevelExpression() *Function {
	expression := []*ExprNode{parser.ParseExpression()}
	prototype := &FunctionProtoype{}

	return &Function{
		Protoype: prototype,
		Body:     expression,
	}
}

// parses an (arbitrarily long) expression
func (parser *Parser) ParseExpression() *ExprNode {
	if parser.ShowDebug {
		fmt.Println("parsing expression...")
	}
	LHS := parser.ParsePrimaryExpression()

	// parse RHS, if it exists...
	return parser.ParseBinaryOpRHS(0, LHS)
}

// Parses RHS of a binary expression.
// Operator-Precedence Parsing Algorithm
func (parser *Parser) ParseBinaryOpRHS(current_precedence int, LHS *ExprNode) *ExprNode {
	for { // recursively parse RHS
		cur_tok := parser.CurrentToken()
		precedence := parser.OperatorPrecedence(cur_tok.Value)

		if precedence < current_precedence {
			fmt.Println("finished parsing binary expr...")
			return LHS
		}

		// so we must have a binary Operator
		binop := cur_tok
		parser.Advance()

		RHS := parser.ParsePrimaryExpression()

		next_precedence := parser.OperatorPrecedence(parser.CurrentToken().Value)
		if precedence < next_precedence {
			RHS = parser.ParseBinaryOpRHS(precedence+1, RHS)
		}

		// merge LHS and RHS
		LHS = &ExprNode{
			Expr: &BinaryExprNode{Operator: binop.Value, LHS: LHS, RHS: RHS},
		}

	}
}

// parses a parenexpr
// <paren_expr> ::= ( expr )
func (parser *Parser) ParseParenExpression() *ExprNode {
	cur_tok := parser.CurrentToken()
	if cur_tok.ID != tokens.OPEN_PAREN {
		parser.ParserError("ParseParenExpression", "(", cur_tok.Repr(), cur_tok.LineNumber)
	}
	parser.Advance()

	// get expression
	expr := parser.ParseExpression()

	cur_tok = parser.CurrentToken()
	if cur_tok.ID != tokens.CLOSE_PAREN {
		parser.ParserError("ParseParenExpression", ")", cur_tok.Repr(), cur_tok.LineNumber)
	}
	parser.Advance()

	return expr
}
