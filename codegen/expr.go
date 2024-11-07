package codegen

import (
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func DoubleCodeGen(double_literal parser.DoubleLiteral) constant.Constant {
	return constant.NewFloat(types.Double, double_literal.Value)
}

func (ctx *Context) VariableReferenceCodeGen(var_ref parser.VariableReference) value.Value {
	ident_v := ctx.lookupVariable(var_ref.Identifier)
	return ctx.NewLoad(types.Double, ident_v)
}

func (ctx *Context) FunctionCallCodeGen(mod *ir.Module, function_call parser.FunctionCall) value.Value {
	call := ctx.NewCall(GetFunc(mod, function_call.FunctionName))
	for _, arg := range function_call.Args {
		// add arguments
		call.Args = append(call.Args, ir.NewArg(ctx.ExprCodeGen(*arg)))
	}
	return call
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