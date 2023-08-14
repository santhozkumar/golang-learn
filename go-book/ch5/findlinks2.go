package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        links, err := findLinks(url)

        if err != nil {
            fmt.Println(err)
        }

    for _, link := range links {
        fmt.Println(link)
    }


    }

}


func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("getting %s: %d\n", url, resp.StatusCode)
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("parsing %s as html: %v", url, err)
    }
    return visit(nil, doc), nil
    // return []string{"", ""}, errors.New("got error")
}

func visit(links []string, n *html.Node) []string {
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

// func CountWordsAndImages(url string) (words, images int, err error) {
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return 
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        err = fmt.Errorf("getting %s: %d\n", url, resp.StatusCode)
        return 
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        err = fmt.Errorf("parsing %s as html: %v", url, err)
        return 
    }
    words, images = countWordsAndImages(doc)
    return 
}


func countWordsAndImages(n *html.Node) (words, images int) {

    return 0, 0

}
