package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)


func Start(in io.Reader, out io.Writer){
    const PROMPT = ">> "
    fmt.Println("starting read")

    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return 
        }
        line := scanner.Text()
        if line == "exit" {
            return 
        }

        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
            fmt.Printf("%+v",tok)

        }
    }

}

