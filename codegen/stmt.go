package codegen

import (
	"fmt"
	"reflect"
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) StmtCodeGen(mod *ir.Module, stmt parser.Stmt, f *ir.Func) {
	fmt.Println(reflect.TypeOf(stmt))
	switch stmt := stmt.(type) {
	case *parser.StmtDefine:
		fmt.Println("stmt define found")
		ctx.StmtDefineCodeGen(mod, *stmt, f)
	case *parser.StmtAssign:
		ctx.StmtAssignCodeGen(mod, *stmt, f)
	case *parser.StmtReturn:
		ctx.ReturnExprCodeGen(mod, *stmt, f)
	case *parser.StmtFunctionCall:
		ctx.StmtFunctionCallCodeGen(mod, *stmt, f)
	case *parser.StmtIfThen:
		ctx.IfThenCodeGen(mod, *stmt, f)
	}
}

func (ctx *Context) StmtDefineCodeGen(mod *ir.Module, stmt_define parser.StmtDefine, f *ir.Func) {
	fmt.Println("CODEGENING A DEFINE STMT...")
	// initialize it to the expr or to 0.
	// first, let's allocate space
	v := ctx.NewAlloca(stmt_define.Typ)
	if stmt_define.HasExpr {
		// store the expr
		ctx.NewStore(ctx.ExprCodeGen(stmt_define.Expr), v)
	} else {
		// store a 0
		ctx.NewStore(constant.NewFloat(types.Double, 0), v)
	}
	fmt.Println(stmt_define.Identifier)
	ctx.vars[stmt_define.Identifier] = v
}

func (ctx *Context) StmtAssignCodeGen(mod *ir.Module, stmt_assign parser.StmtAssign, f *ir.Func) {
	v := ctx.lookupVariable(stmt_assign.Identifier)
	ctx.NewStore(ctx.ExprCodeGen(stmt_assign.Expr), v)
}

func (ctx *Context) ReturnExprCodeGen(mod *ir.Module, stmt_ret parser.StmtReturn, f *ir.Func) {
	fmt.Println("CODEGENING A RETURN STMT...")
	ctx.NewRet(ctx.ExprCodeGen(stmt_ret.ReturnExpr))
}

func (ctx *Context) IfThenCodeGen(mod *ir.Module, stmt_if_then parser.StmtIfThen, f *ir.Func) {
	my_idx := *ctx.if_idx
	*ctx.if_idx += 1

	// create new then block
	then_ctx := ctx.NewContext(f.NewBlock(fmt.Sprintf("if.then%d", my_idx)))
	for _, stmt := range stmt_if_then.Then {
		then_ctx.StmtCodeGen(mod, stmt, f)
	}

	else_ctx := ctx.NewContext(f.NewBlock(fmt.Sprintf("if.else%d", my_idx)))
	if stmt_if_then.Else != nil {
		// create a new else block
		for _, stmt := range stmt_if_then.Else {
			else_ctx.StmtCodeGen(mod, stmt, f)
		}
	}

	ctx.NewCondBr(ctx.ExprCodeGen(stmt_if_then.Condition), then_ctx.Block, else_ctx.Block)
	
	// make sure we skip the else
	if !then_ctx.HasTerminator() {
		leave_if := f.NewBlock(fmt.Sprintf("leave.if%d", my_idx))
		then_ctx.NewBr(leave_if)
	}
}  

func StmtFunctionDeclarationCodeGen(mod *ir.Module, stmt_function_decl parser.StmtFunctionDeclaration) {
	// create function
	if stmt_function_decl.Function.Protoype.FunctionName == "main" {
		fmt.Println("MAIN FUNCTION FOUND!!!")
	} else {
		fmt.Println(stmt_function_decl.Function.Protoype.FunctionName)
	}
	f := mod.NewFunc(stmt_function_decl.Function.Protoype.FunctionName, types.Double)
	entry := f.NewBlock("entry")
	ctx := NewContext(entry)
	for _, arg := range stmt_function_decl.Function.Protoype.Args {
		// add arguments
		param := ir.NewParam(arg, types.Double)
		f.Params = append(f.Params, param)
		// add variable to context so we know its defined
		location := ctx.NewAlloca(types.Double)
		ctx.NewStore(param, location)
		ctx.vars[arg] = location
	}
	// generate body of function
	fmt.Println(len(stmt_function_decl.Function.Body))
	for _, stmt := range stmt_function_decl.Function.Body {
		fmt.Println("trying to codegen a stmt...")
		ctx.StmtCodeGen(mod, stmt, f)
	}

	if !ctx.HasTerminator() {
		ctx.NewRet(constant.NewFloat(types.Double, 0))
	}
}

func (ctx *Context) StmtFunctionCallCodeGen(mod *ir.Module, stmt_function_call parser.StmtFunctionCall, f *ir.Func) value.Value {
	call := ctx.NewCall(GetFunc(mod, stmt_function_call.FunctionCall.FunctionName))
	for _, arg := range stmt_function_call.FunctionCall.Args {
		// add arguments
		call.Args = append(call.Args, ir.NewArg(ctx.ExprCodeGen(*arg)))
	}
	return call
}