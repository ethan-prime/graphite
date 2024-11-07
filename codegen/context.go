package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"fmt"
)

type Context struct {
	*ir.Block
	parent *Context
	vars   map[string]value.Value
	if_idx int
	loop_idx int
}

func NewContext(b *ir.Block) *Context {
	return &Context{
		Block: b,
		parent:   nil,
		vars:     make(map[string]value.Value),
		if_idx: 0,
		loop_idx: 0,
	}
}

func (c *Context) NewContext(b *ir.Block) *Context {
	ctx := NewContext(b)
	ctx.parent = c
	ctx.if_idx, ctx.loop_idx = c.if_idx, c.loop_idx
	return ctx
}

func (c Context) lookupVariable(name string) value.Value {
	if v, ok := c.vars[name]; ok {
		return v
	} else if c.parent != nil {
		return c.parent.lookupVariable(name)
	} else {
		fmt.Printf("variable: `%s`\n", name)
		panic("no such variable")
	}
}

func (c *Context) HasTerminator() bool {
	if c.Term != nil {
		return true
	}
	return false
}