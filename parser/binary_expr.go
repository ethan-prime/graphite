package parser

var BinaryOpPrecedence = map[string]int{
	"<": 10,
	"+": 20,
	"-": 20,
	"*": 40,
	"/": 40,
}

type BinaryExprNode struct {
	operator string
	LHS *ExprNode
	RHS *ExprNode
}

// retrieves the precedence of a binary operator, -1 if it doesn't exist!
func (parser *Parser) OperatorPrecedence(binop string) int {
	precedence, ok := BinaryOpPrecedence[binop]
	if !ok {
		return -1
	}
	return precedence
}