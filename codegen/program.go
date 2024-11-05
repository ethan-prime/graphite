package codegen

import (
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) ProgramCodeGen(program_node parser.ProgramNode) value.Value {
	for _, stmt := range program_node.Stmts {
		ctx.StmtCodeGen(stmt)
	}
	panic("")
}