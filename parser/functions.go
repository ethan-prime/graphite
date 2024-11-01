package parser

import (
	"fmt"

	"github.com/ethan-prime/graphite/tokens"
)

type FunctionCall struct {
	FunctionName string
	Args         []*ExprNode
}

type FunctionProtoype struct { // includes function name and parameters/args.
	FunctionName string
	Args         []string
}

type Function struct {
	Protoype *FunctionProtoype
	Body     []*ExprNode
}

// <function_prototype> ::= <identifier> ( <arguments> )
func (parser *Parser) ParseFunctionPrototype() *FunctionProtoype {
	if parser.ShowDebug {
		fmt.Println("parsing function prototype...")
	}
	cur_tok := parser.CurrentToken()
	if cur_tok.ID != tokens.IDENTIFIER {
		parser.ParserError("ParseFunctionPrototype", "Identifier", cur_tok.Repr(), cur_tok.LineNumber)
	}

	identifier := cur_tok.Value

	parser.Advance()
	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.OPEN_PAREN {
		parser.ParserError("ParseFunctionPrototype", "(", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance() // eat the (

	var args []string

	for cur_tok := parser.CurrentToken(); cur_tok.ID == tokens.IDENTIFIER; cur_tok = parser.CurrentToken() {
		// we have an argument to add
		args = append(args, cur_tok.Value)
		parser.Advance()
		// we should have a comma now if there is another argument
		if parser.PeekToken().ID == tokens.IDENTIFIER {
			if parser.CurrentToken().ID != tokens.COMMA {
				parser.ParserError("ParseFunctionPrototype", ",", parser.CurrentToken().Value, parser.CurrentToken().LineNumber)
			}
			parser.Advance() // eat the comma
		}
	}

	if parser.CurrentToken().ID != tokens.CLOSE_PAREN {
		parser.ParserError("ParseFunctionPrototype", ")", parser.CurrentToken().Value, parser.CurrentToken().LineNumber)
	}

	parser.Advance() // eat )

	// return the Prototype struct pointer
	return &FunctionProtoype{
		FunctionName: identifier,
		Args:         args,
	}
}

// <function_definition> ::= def <function_prototype> { <function_body> }
func (parser *Parser) ParseFunctionDeclaration() *Function {
	if parser.ShowDebug {
		fmt.Println("parsing function declaration...")
	}
	cur_tok := parser.CurrentToken()
	if cur_tok.ID != tokens.KEYW_DEF {
		parser.ParserError("ParseFunctionDeclaration", "def", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	prototype := parser.ParseFunctionPrototype()
	var body []*ExprNode

	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.OPEN_BRACE {
		parser.ParserError("ParseFunctionPrototype", "{", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	for cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE; cur_tok = parser.CurrentToken() {
		// we have an expr to add to the function body
		body = append(body, parser.ParseExpression())
	}

	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE {
		parser.ParserError("ParseFunctionPrototype", "}", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	return &Function{
		Protoype: prototype,
		Body:     body,
	}
}
