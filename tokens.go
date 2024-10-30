package main

type TokenID int

// TOKEN IDS
const (
	EOF TokenID = iota
	DEF // for defining functions
	IDENTIFIER
	DOUBLE
	RET // return
	UNKNOWN
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
	case DEF:
		return "def"
	case IDENTIFIER:
		return "identifier"
	case DOUBLE:
		return "double"
	case RET:
		return "ret"
	default:
		return "**UNKNOWN TOKEN**"
	}
}