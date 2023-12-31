package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/israel-duff/go-monkey-interpreter/pkg/lexer"
	"github.com/israel-duff/go-monkey-interpreter/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
