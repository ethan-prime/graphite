package codegen

import (
	"github.com/ethan-prime/graphite/parser"
	"github.com/llir/llvm/ir"
	"fmt"
)

func ProgramCodeGen(program_node parser.ProgramNode) {
	mod := ir.NewModule()

	fmt.Println(len(program_node.Stmts))

	for _, stmt := range program_node.Stmts {
		switch stmt := stmt.(type) {
		case *parser.StmtFunctionDeclaration:
			StmtFunctionDeclarationCodeGen(mod, *stmt)
		}
	}

	// make sure the main function is defined
	GetFunc(mod, "main")
	fmt.Println(mod.String())
}