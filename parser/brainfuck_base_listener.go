// Generated from src/github.com/gonccalo/GOfuck/Brainfuck.g4 by ANTLR 4.6.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBrainfuckListener is a complete listener for a parse tree produced by BrainfuckParser.
type BaseBrainfuckListener struct{}

var _ BrainfuckListener = &BaseBrainfuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBrainfuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBrainfuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBrainfuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBrainfuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBrainfuckListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBrainfuckListener) ExitProgram(ctx *ProgramContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseBrainfuckListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseBrainfuckListener) ExitCommand(ctx *CommandContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseBrainfuckListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseBrainfuckListener) ExitLoop(ctx *LoopContext) {}

// EnterOp is called when production op is entered.
func (s *BaseBrainfuckListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseBrainfuckListener) ExitOp(ctx *OpContext) {}
