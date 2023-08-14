// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"
    "strings"

	"golang.org/x/net/html"
)

//!+
func mainw() {
    s := `<p>Links:</p><ul><li>
    <a href="foo">Foo</a>
    <li><a href="/bar/baz">BarBaz</a>
    </ul>
    // <div><a href="www.google.com">google</a></div>`
    doc, err := html.Parse(strings.NewReader(s))
	// doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "summary: %v\n", err)
		os.Exit(1)
	}
	var mapping = make(map[string]int)
	summary(mapping, doc)

	// summary
	for k, v := range mapping {
		fmt.Println(k, v)
	}
}

func summary(mapping map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		mapping[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		summary(mapping, c)
	}
}
