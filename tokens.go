package main

type TokenID int

// TOKEN IDS
const (
	EOF TokenID = iota
	KEYW_DEF // for defining functions
	IDENTIFIER
	DOUBLE
	KEYW_RET // return
	UNKNOWN
	OPEN_PAREN // (
	CLOSE_PAREN // )
	OPEN_BRACE // {
	CLOSE_BRACE // }
	ARROW // =>
	KEYW_DBL // dbl
	EQUAL // =
	PLUS // +
	MINUS // -
	ASTERIK // *
	SLASH // /
	COMMA // ,
)

type Token struct {
	id TokenID // unique id of token
	value string // can hold values such as ints, doubles, identifiers, etc.
}

// repr of a token. example: def -> "def"
func (t Token) Repr() string {
	switch t.id {
	case EOF:
		return "eof"
	case KEYW_DEF:
		return "def"
	case IDENTIFIER:
		return "identifier"
	case DOUBLE:
		return "double"
	case KEYW_RET:
		return "ret"
	case OPEN_PAREN:
		return "("
	case CLOSE_PAREN:
		return ")"
	case OPEN_BRACE:
		return "{"
	case CLOSE_BRACE:
		return "}"
	case ARROW:
		return "=>"
	case KEYW_DBL:
		return "(keyw) dbl"
	case EQUAL:
		return "="
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case ASTERIK:
		return "*"
	case SLASH:
		return "/"
	case COMMA:
		return ","
	default:
		return "**UNKNOWN TOKEN**"
	}
}