package codegen

import (
	"fmt"
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) StmtCodeGen(stmt parser.Stmt) value.Value {
	panic("")
}

func (ctx *Context) StmtDefineCodeGen(stmt_define parser.StmtDefine) {
	// initialize it to the expr or to 0.
	// first, let's allocate space
	v := ctx.NewAlloca(stmt_define.Typ)
	if stmt_define.Expr == nil {
		// store a 0
		ctx.NewStore(constant.NewFloat(types.Double, 0), v)
	} else {
		// store the expr
		ctx.NewStore(ctx.ExprCodeGen(stmt_define.Expr.(parser.ExprNode)), v)
	}
	ctx.vars[stmt_define.Identifier] = v
}

func (ctx *Context) StmtAssignCodeGen(stmt_assign parser.StmtAssign) {
	v := ctx.lookupVariable(stmt_assign.Identifier)
	if stmt_assign.Expr == nil {
		panic("[ graphite compiler ] no expresssion provided for assign!")
	}
	ctx.NewStore(ctx.ExprCodeGen(stmt_assign.Expr.(parser.ExprNode)), v)
}

func (ctx *Context) ReturnExprCodeGen(stmt_ret parser.StmtReturn) {
	if stmt_ret.ReturnExpr == nil {
		// return 0.0
		ctx.NewRet(constant.NewFloat(types.Double, 0))
	} else {
		ctx.NewRet(ctx.ExprCodeGen(stmt_ret.ReturnExpr.(parser.ExprNode)))
	}
}

func (ctx *Context) IfThenCodeGen(stmt_if_then parser.StmtIfThen, f *ir.Func) {
	my_idx := ctx.if_idx
	ctx.if_idx += 1

	// create new then block
	then_ctx := ctx.NewContext(f.NewBlock(fmt.Sprintf("if.then%d", my_idx)))
	for _, stmt := range stmt_if_then.Then {
		then_ctx.StmtCodeGen(stmt)
	}

	else_ctx := ctx.NewContext(f.NewBlock(fmt.Sprintf("if.else%d", my_idx)))
	if stmt_if_then.Else != nil {
		// create a new else block
		for _, stmt := range stmt_if_then.Else {
			else_ctx.StmtCodeGen(stmt)
		}
	}

	ctx.NewCondBr(ctx.ExprCodeGen(stmt_if_then.Condition.(parser.ExprNode)), then_ctx.Block, else_ctx.Block)
	
	// make sure we skip the else
	if !then_ctx.HasTerminator() {
		leave_if := f.NewBlock(fmt.Sprintf("leave.if%d", my_idx))
		then_ctx.NewBr(leave_if)
	}
	
}  