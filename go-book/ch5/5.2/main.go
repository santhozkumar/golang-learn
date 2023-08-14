package main


import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func main(){

    
    var counts  = make(map[string]int)
    // counts = make(map[string]int)

    s := `<p>Links:</p><ul><li>
    <a href="foo">Foo</a>
    <li><a href="/bar/baz">BarBaz</a>
    </ul>
    <div><a href="www.google.com">google</a></div>`
    // s := "<html></html>"
    doc, err := html.Parse(strings.NewReader(s))
    // doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
        os.Exit(1)
    }

    findCounts(counts, doc)

    for k,v := range counts {
        fmt.Printf("%s: %d\n", k, v)
    }

}


func findCounts(counts map[string]int, n *html.Node) {
    if n.Type == html.ElementNode {
        counts[n.Data]++
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        findCounts(counts, c)
    }

}




