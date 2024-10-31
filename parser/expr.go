package parser

// define base expr for all other exprs
type Expr interface {
	ParseExpr()
}

type ExprNode struct {
	Expr Expr
}