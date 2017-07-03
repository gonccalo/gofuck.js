// Generated from parser/Brainfuck.g4 by ANTLR 4.6.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBrainfuckVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBrainfuckVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBrainfuckVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBrainfuckVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBrainfuckVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}
