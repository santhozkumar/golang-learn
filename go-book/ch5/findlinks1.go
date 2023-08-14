package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)



func main2(){
    s := `<p>Links:</p><ul><li>
    <a href="foo">Foo</a>
    <li><a href="/bar/baz">BarBaz</a>
    </ul>
    <div><a href="www.google.com">google</a></div>`
    doc, err := html.Parse(strings.NewReader(s))
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
        os.Exit(1)
    }

    for _, link := range visit(nil, doc){
        fmt.Println(link)
    }
}


func visit2(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a"{
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                links = append(links, attr.Val)
            }
       }
   }

    for c:=n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
    
}
