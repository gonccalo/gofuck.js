package main

import(
	"github.com/gonccalo/GOfuck/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"os"
)

type TreeShapeListener struct {
	*parser.BaseBrainfuckListener
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	inFile, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBrainfuckLexer(inFile)
	stream := antlr.NewCommonTokenStream(lexer,0)
	pars := parser.NewBrainfuckParser(stream)
	pars.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	pars.BuildParseTrees = true
	tree := pars.Program()
	antlr.ParseTreeWalkerDefault.Walk(new(TreeShapeListener), tree)
}
