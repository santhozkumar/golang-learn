package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


func main(){
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
        os.Exit(1)
    }
}
