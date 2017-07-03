// Generated from parser/Brainfuck.g4 by ANTLR 4.6.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BrainfuckListener is a complete listener for a parse tree produced by BrainfuckParser.
type BrainfuckListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)
}
