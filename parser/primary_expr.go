package parser

import (
	"log"
	"github.com/ethan-prime/graphite/tokens"
	"strconv"
)

type PrimaryExpr interface {}

type PrimaryExprNode struct {
	PrimaryExpr PrimaryExpr
}

type DoubleLiteral struct {
	Value float64
}

type VariableReference struct {
	Identifier string
}

// parses a primary expression
// <primary_expr> ::= <double> | <identifer_expr>
func (parser *Parser) ParsePrimaryExpression() *ExprNode {
	cur_tok := parser.CurrentToken()
	switch cur_tok.ID {
	case tokens.DOUBLE:
		return parser.ParseDoubleExpression()
	case tokens.IDENTIFIER:
		return parser.ParseIdentifierExpression()
	default:
		parser.ParserError("ParsePrimaryExpression", "Primary Expression", cur_tok.Repr(), cur_tok.LineNumber)
	}
	return &ExprNode{}
}

// parses a double literal
func (parser *Parser) ParseDoubleExpression() *ExprNode {
	cur_tok := parser.CurrentToken()

	if cur_tok.ID != tokens.DOUBLE {
		parser.ParserError("ParseDoubleExpression", "Double", cur_tok.Repr(), cur_tok.LineNumber)
	}

	f, err := strconv.ParseFloat(cur_tok.Value, 64)
	if err != nil {
		log.Fatal(err)
	}
	
	expr := &ExprNode{
		Expr: &DoubleLiteral{Value: f},
	}

	return expr
}

// parses an indentifier_expr
// <identifer_expr> ::= <identifier> | <identifier> ( <parameters> )
func (parser *Parser) ParseIdentifierExpression() *ExprNode {
	cur_tok := parser.CurrentToken()

	if cur_tok.ID != tokens.IDENTIFIER {
		parser.ParserError("ParseIdentifierExpression", "Identifier", cur_tok.Repr(), cur_tok.LineNumber)
	}

	identifier := cur_tok.Value
	expr := &ExprNode{}
	var args []*ExprNode
	parser.Advance()

	if cur_tok.ID != tokens.OPEN_PAREN {
		// we just have a variable access
		expr.Expr = &VariableReference{Identifier: identifier}
		return expr
	}
	
	parser.Advance()
	cur_tok = parser.CurrentToken()

	if cur_tok.ID != tokens.CLOSE_PAREN {
		for {
			arg := parser.ParseExpression()
			args = append(args, arg)

			cur_tok = parser.CurrentToken()

			if cur_tok.ID == tokens.CLOSE_PAREN {
				break
			}

			if cur_tok.ID != tokens.COMMA {
				parser.ParserError("ParseIdentifierExpression", ",", cur_tok.Repr(), cur_tok.LineNumber)
			}
			parser.Advance() // eat the comma to parse next argument
		}
	}

	parser.Advance() // eat the )

	expr.Expr = FunctionCall{FunctionName: identifier, Args: args}
	return expr
}