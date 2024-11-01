package parser

type FunctionCall struct  {
	FunctionName string
	Args []*ExprNode
}

type FunctionProtoype struct { // includes function name and parameters/args.
	FunctionName string
	Args []*ExprNode
}

type Function struct {
	Protoype FunctionProtoype
	Body *ExprNode
}