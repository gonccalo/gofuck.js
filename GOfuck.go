package main

import(
	"github.com/gonccalo/GOfuck/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"flag"
	//"unicode/utf8"
)

type GoFuckListener struct {
	*parser.BaseBrainfuckListener
}

func (this *GoFuckListener) EnterProgram(ctx *parser.ProgramContext){
	fmt.Printf("var mem 	= new Array(30000).fill(0);\n")
	fmt.Printf("var pointer = 0;\n")
	fmt.Printf("var outDiv  = document.getElementById(\"output\");\n")
	fmt.Printf("var enter   = document.getElementById(\"enter\");\n")
	fmt.Printf("var input   = document.getElementById(\"input\");\n")
	fmt.Printf("var info    = document.getElementById(\"info\");\n")
	fmt.Printf("var gen 	= main();\n")
	fmt.Printf("gen.next();\n")
	fmt.Printf("enter.onclick = function() {\n")
	fmt.Printf("    mem[pointer] = input.value.charCodeAt(0);\n")
	fmt.Printf("    gen.next();\n")
	fmt.Printf("};\n")

	fmt.Printf("function* main() {\n")
}

func (this *GoFuckListener) EnterOp(ctx *parser.OpContext){
	switch ctx.GetCmd().GetTokenType() {
	case parser.BrainfuckParserTAPE_INC:
		fmt.Println("mem[pointer]++;")
	case parser.BrainfuckParserTAPE_DEC:
		fmt.Println("mem[pointer]--;")
	case parser.BrainfuckParserTAPE_LEFT:
		fmt.Println("pointer--;")
	case parser.BrainfuckParserTAPE_RIGHT:
		fmt.Println("pointer++;")
	case parser.BrainfuckParserINPUT:
		/*behold the power of async input
		fmt.Printf("enter.onclick = function() {\n")
		fmt.Printf("mem[pointer] = input.value.charCodeAt(0);\n")
		toClose++
		*/
		fmt.Printf("enter.disabled = false;\n")
		fmt.Printf("info.innerText = \"Input text now\";\n")
		fmt.Printf("yield;\n")
		fmt.Printf("info.innerText = \"\";\n")
		fmt.Printf("enter.disabled = true;\n")
	case parser.BrainfuckParserOUTPUT:
		fmt.Printf("outDiv.innerText = outDiv.innerText + String.fromCharCode(mem[pointer]);\n")
		fmt.Printf("console.log(String.fromCharCode(mem[pointer]));\n")
	}
}

func (this *GoFuckListener) EnterLoop(ctx *parser.LoopContext){
	fmt.Printf("while (mem[pointer] != 0) {\n")
}

func (this *GoFuckListener) ExitLoop(ctx *parser.LoopContext){
	fmt.Printf("}\n")
}

func (this *GoFuckListener) ExitProgram(ctx *parser.ProgramContext){
	/*for i := 0; i < toClose; i++ {
		fmt.Printf("}\n")	
	}*/
	//exit main
	fmt.Printf("}\n")
}

var filename string
var size int
var toClose int

func init() {
	//fuck nodejs it can't even do sync reads from stdin wtf
	//flag.BoolVar(&useNodejs, "nodejs", false, "Set to true to generate nodejs code")
	flag.IntVar(&size, "mem", 30000, "Set the memory size")
	flag.StringVar(&filename, "i", "", "File to run")
	flag.Parse()
}

func main() {
	inFile, _ := antlr.NewFileStream(filename)
	lexer := parser.NewBrainfuckLexer(inFile)
	stream := antlr.NewCommonTokenStream(lexer,0)
	pars := parser.NewBrainfuckParser(stream)
	pars.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	pars.BuildParseTrees = true
	tree := pars.Program()
	antlr.ParseTreeWalkerDefault.Walk(new(GoFuckListener), tree)
}
	