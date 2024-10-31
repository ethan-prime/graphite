package parser

type BinaryExprNode struct {
	operator string
	LHS ExprNode
	RHS ExprNode
}