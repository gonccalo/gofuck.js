// Generated from parser/Brainfuck.g4 by ANTLR 4.6.

package parser // Brainfuck

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 11, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 3, 3, 3, 5, 3, 19, 10, 3, 3, 4, 3, 4, 7, 4, 23,
	10, 4, 12, 4, 14, 4, 26, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 2, 2, 6,
	2, 4, 6, 8, 2, 3, 3, 2, 5, 10, 30, 2, 13, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2,
	6, 20, 3, 2, 2, 2, 8, 29, 3, 2, 2, 2, 10, 12, 5, 4, 3, 2, 11, 10, 3, 2,
	2, 2, 12, 15, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 3,
	3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 16, 19, 5, 6, 4, 2, 17, 19, 5, 8, 5, 2,
	18, 16, 3, 2, 2, 2, 18, 17, 3, 2, 2, 2, 19, 5, 3, 2, 2, 2, 20, 24, 7, 3,
	2, 2, 21, 23, 5, 4, 3, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22,
	3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2,
	27, 28, 7, 4, 2, 2, 28, 7, 3, 2, 2, 2, 29, 30, 9, 2, 2, 2, 30, 9, 3, 2,
	2, 2, 5, 13, 18, 24,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'+'", "'-'", "'<'", "'>'", "','", "'.'",
}

var symbolicNames = []string{
	"", "", "", "TAPE_INC", "TAPE_DEC", "TAPE_LEFT", "TAPE_RIGHT", "INPUT",
	"OUTPUT", "END",
}

var ruleNames = []string{
	"program", "command", "loop", "op",
}

type BrainfuckParser struct {
	*antlr.BaseParser
}

func NewBrainfuckParser(input antlr.TokenStream) *BrainfuckParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(BrainfuckParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Brainfuck.g4"

	return this
}

// BrainfuckParser tokens.
const (
	BrainfuckParserEOF        = antlr.TokenEOF
	BrainfuckParserT__0       = 1
	BrainfuckParserT__1       = 2
	BrainfuckParserTAPE_INC   = 3
	BrainfuckParserTAPE_DEC   = 4
	BrainfuckParserTAPE_LEFT  = 5
	BrainfuckParserTAPE_RIGHT = 6
	BrainfuckParserINPUT      = 7
	BrainfuckParserOUTPUT     = 8
	BrainfuckParserEND        = 9
)

// BrainfuckParser rules.
const (
	BrainfuckParserRULE_program = 0
	BrainfuckParserRULE_command = 1
	BrainfuckParserRULE_loop    = 2
	BrainfuckParserRULE_op      = 3
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllCommand() []ICommandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandContext)(nil)).Elem())
	var tst = make([]ICommandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandContext)
		}
	}

	return tst
}

func (s *ProgramContext) Command(i int) ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BrainfuckVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BrainfuckParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BrainfuckParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserT__0)|(1<<BrainfuckParserTAPE_INC)|(1<<BrainfuckParserTAPE_DEC)|(1<<BrainfuckParserTAPE_LEFT)|(1<<BrainfuckParserTAPE_RIGHT)|(1<<BrainfuckParserINPUT)|(1<<BrainfuckParserOUTPUT))) != 0 {
		{
			p.SetState(8)
			p.Command()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *CommandContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BrainfuckVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BrainfuckParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BrainfuckParserRULE_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BrainfuckParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Loop()
		}

	case BrainfuckParserTAPE_INC, BrainfuckParserTAPE_DEC, BrainfuckParserTAPE_LEFT, BrainfuckParserTAPE_RIGHT, BrainfuckParserINPUT, BrainfuckParserOUTPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Op()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) AllCommand() []ICommandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandContext)(nil)).Elem())
	var tst = make([]ICommandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandContext)
		}
	}

	return tst
}

func (s *LoopContext) Command(i int) ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BrainfuckVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BrainfuckParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BrainfuckParserRULE_loop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(BrainfuckParserT__0)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserT__0)|(1<<BrainfuckParserTAPE_INC)|(1<<BrainfuckParserTAPE_DEC)|(1<<BrainfuckParserTAPE_LEFT)|(1<<BrainfuckParserTAPE_RIGHT)|(1<<BrainfuckParserINPUT)|(1<<BrainfuckParserOUTPUT))) != 0 {
		{
			p.SetState(19)
			p.Command()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(BrainfuckParserT__1)
	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCmd returns the cmd token.
	GetCmd() antlr.Token

	// SetCmd sets the cmd token.
	SetCmd(antlr.Token)

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cmd    antlr.Token
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) GetCmd() antlr.Token { return s.cmd }

func (s *OpContext) SetCmd(v antlr.Token) { s.cmd = v }

func (s *OpContext) TAPE_INC() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserTAPE_INC, 0)
}

func (s *OpContext) TAPE_DEC() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserTAPE_DEC, 0)
}

func (s *OpContext) TAPE_LEFT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserTAPE_LEFT, 0)
}

func (s *OpContext) TAPE_RIGHT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserTAPE_RIGHT, 0)
}

func (s *OpContext) INPUT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserINPUT, 0)
}

func (s *OpContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserOUTPUT, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitOp(s)
	}
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BrainfuckVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BrainfuckParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BrainfuckParserRULE_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*OpContext).cmd = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserTAPE_INC)|(1<<BrainfuckParserTAPE_DEC)|(1<<BrainfuckParserTAPE_LEFT)|(1<<BrainfuckParserTAPE_RIGHT)|(1<<BrainfuckParserINPUT)|(1<<BrainfuckParserOUTPUT))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*OpContext).cmd = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
