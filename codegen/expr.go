package codegen

import (
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

func DoubleCodeGen(double_literal parser.DoubleLiteral) constant.Constant {
	return constant.NewFloat(types.Double, double_literal.Value)
}

func (ctx *Context) VariableReferenceCodeGen(var_ref parser.VariableReference) value.Value {
	return ctx.lookupVariable(var_ref.Identifier)
}

func (ctx *Context) ExprCodeGen(expr_node parser.ExprNode) value.Value {
	switch expr_node := expr_node.Expr.(type) {
	case *parser.DoubleLiteral:
		return DoubleCodeGen(*expr_node)
	case *parser.VariableReference:
		return ctx.VariableReferenceCodeGen(*expr_node)
	case *parser.BinaryExprNode:
		l, r := ctx.ExprCodeGen(*expr_node.LHS), ctx.ExprCodeGen(*expr_node.RHS)
		if expr_node.Operator == "+" {
			return ctx.NewAdd(l, r)
		} else if expr_node.Operator == "-" {
			return ctx.NewSub(l, r)
		} else if expr_node.Operator == "*" {
			return ctx.NewMul(l, r)
		} else if expr_node.Operator == "/" {
			return ctx.NewFDiv(l, r)
		}
	}
	panic("[ graphite compiler ] ExprCodeGen(): unknown expression!")
}