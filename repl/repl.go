package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/satoshi03/go-monkey-interpreter/evaluator"
	"github.com/satoshi03/go-monkey-interpreter/lexer"
	"github.com/satoshi03/go-monkey-interpreter/parser"
)

const PROMPT = ">> "
const MONKEY = `
w  c(..)o   (
 \__(-)    __)
     /\   (
    /(_)___)
   w /|
    | \
    m  m
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, MONKEY)
		io.WriteString(out, "Woops! We ran into some monkey business here!\n")
		io.WriteString(out, " parser errors:\n")
		io.WriteString(out, "\t"+msg+"\n")
	}
}
