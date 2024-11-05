package parser

import (
	"fmt"
	"github.com/llir/llvm/ir/types"
	"github.com/ethan-prime/graphite/tokens"
)

type Stmt interface{}

type StmtDefine struct {
	Stmt
	Identifier string
	Typ        types.Type
	Expr Expr
}

type StmtAssign struct {
	Stmt
	Identifier string
	Expr Expr
}

type StmtReturn struct {
	Stmt
	ReturnExpr Expr
}

type StmtFunctionCall struct {
	Stmt
	FunctionCall FunctionCall
}

type StmtFunctionDeclaration struct {
	Stmt
	Function Function
}

type StmtExpression struct {
	Stmt
	Expr Expr
}

type StmtIfThen struct {
	Stmt
	Condition Expr
	Then Stmt
	Else Stmt
}

func (parser *Parser) ParseStatement() Stmt {
	switch parser.CurrentToken().ID {
	case tokens.KEYW_DEF:
		f := parser.ParseFunctionDeclaration()
		stmt_function_decl := StmtFunctionDeclaration{Function: *f}
		fmt.Println("parsed function declaration")
		return stmt_function_decl
	case tokens.SEMICOLON:
		parser.Advance()
	case tokens.KEYW_LET:
		fmt.Println("parsed a variable declaration stmt...")
		return parser.ParseVariableDefinition()
	case tokens.IDENTIFIER:
		if parser.PeekToken().ID == tokens.OPEN_PAREN {
			fmt.Println("parsed identifier expression (probably a function call)...")
			call := parser.ParseIdentifierExpression().Expr
			return StmtFunctionCall{FunctionCall: call.(FunctionCall)}
		} else if parser.PeekToken().ID == tokens.EQUAL {
			fmt.Println("parsed an assignment stmt...")
			return parser.ParseAssignment()
		}
	case tokens.KEYW_RET:
		fmt.Println("parsed a return stmt...")
		return parser.ParseReturn()
	default:
		parser.ParserError("ParseStatement", "a statment", parser.CurrentToken().Repr(), parser.CurrentToken().LineNumber)
		panic("")
	}
	parser.ParserError("ParseStatement", "a statment", parser.CurrentToken().Repr(), parser.CurrentToken().LineNumber)
	panic("")
}