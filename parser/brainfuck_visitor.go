// Generated from parser/Brainfuck.g4 by ANTLR 4.6.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BrainfuckParser.
type BrainfuckVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BrainfuckParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BrainfuckParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by BrainfuckParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by BrainfuckParser#op.
	VisitOp(ctx *OpContext) interface{}
}
