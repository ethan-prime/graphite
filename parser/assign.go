package parser

import (
	"github.com/llir/llvm/ir/types"
	"github.com/ethan-prime/graphite/tokens"
)

// ::= let <identifier> = <expr>
func (parser *Parser) ParseVariableDefinition() *StmtDefine {
	if parser.CurrentToken().ID != tokens.KEYW_LET {
		parser.ParserError("ParseVariableDefinition", "let", parser.CurrentToken().Repr(), parser.CurrentToken().LineNumber)
	}

    parser.Advance()

    if token := parser.CurrentToken(); token.ID != tokens.IDENTIFIER {
		parser.ParserError("ParseVariableDefinition", "variable identifier", token.Repr(), token.LineNumber)
	}

    identifier := parser.CurrentToken().Value

    parser.Advance()

    if token := parser.CurrentToken(); token.ID != tokens.EQUAL {
		return &StmtDefine{
			Identifier: identifier,
			Typ: types.Double,
			Expr: nil,
		}
	}

    parser.Advance()

    expr := parser.ParseExpression()

	return &StmtDefine{
        Identifier: identifier,
        Typ: types.Double,
        Expr: expr,
    }
}

// ::= <identifier> = <expr>
func (parser *Parser) ParseAssignment() *StmtAssign {
    if token := parser.CurrentToken(); token.ID != tokens.IDENTIFIER {
		parser.ParserError("ParseVariableDefinition", "variable identifier", token.Repr(), token.LineNumber)
	}

    identifier := parser.CurrentToken().Value

    parser.Advance()

    if token := parser.CurrentToken(); token.ID != tokens.EQUAL {
		parser.ParserError("ParseVariableDefinition", "=", token.Repr(), token.LineNumber)
	}

    parser.Advance()

    expr := parser.ParseExpression()

	return &StmtAssign{
        Identifier: identifier,
        Expr: expr,
    }
}