package parser

import (
	"github.com/ethan-prime/graphite/tokens"
)

func (parser *Parser) ParseIfThen() *StmtIfThen {
	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.KEYW_IF {
		parser.ParserError("ParseIfthen", "if", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	cond := parser.ParseExpression()

	var then_body []Stmt

	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.OPEN_BRACE {
		parser.ParserError("ParseIfThen", "{", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	for cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE; cur_tok = parser.CurrentToken() {
		// we have an expr to add to the function body
		then_body = append(then_body, parser.ParseStatement())
	}

	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE {
		parser.ParserError("ParseIfThen", "}", cur_tok.Repr(), cur_tok.LineNumber)
	}

	parser.Advance()

	// we need to find out if this is an else or an else if
	if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.KEYW_ELSE {
		// return with a nil else field
		return &StmtIfThen{
			Condition: *cond,
			Then: then_body,
			Else: nil,
			HasElse: false,
		}
	}
	
	var else_body []Stmt

	parser.Advance()
	// we have an else token. lets check if theres an if
	if cur_tok := parser.CurrentToken(); cur_tok.ID == tokens.KEYW_IF {
		// we have an else if !
		else_body = append(else_body, parser.ParseIfThen())
	} else {
		// we should have an open brace
		if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.OPEN_BRACE {
			parser.ParserError("ParseIfThen", "{", cur_tok.Repr(), cur_tok.LineNumber)
		}

		parser.Advance()

		for cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE; cur_tok = parser.CurrentToken() {
			// we have an expr to add to the function body
			else_body = append(else_body, parser.ParseStatement())
		}
	
		if cur_tok := parser.CurrentToken(); cur_tok.ID != tokens.CLOSE_BRACE {
			parser.ParserError("ParseIfThen", "}", cur_tok.Repr(), cur_tok.LineNumber)
		}
	
		parser.Advance()
	}

	return &StmtIfThen{
		Condition: *cond,
		Then: then_body,
		Else: else_body,
		HasElse: true,
	}
}