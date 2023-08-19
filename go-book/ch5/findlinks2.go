package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
    var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
    fmt.Printf("%v %T\n", asciiSpace, asciiSpace)
    for _, url := range os.Args[1:] {
        words, images, err := CountWordsAndImages(url)
        // links, err := findLinks(url)

        if err != nil {
            fmt.Println(err)
        }

    // for _, link := range links {
    //     fmt.Println(link)
    // }

    fmt.Println(words, images)
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

    if n.Type == html.ElementNode {
        if n.Data == "style" || n.Data == "script"{
            return 
        }else if n.Data == "img" {
        // }else if strings.HasPrefix(n.Data, "img"){
            fmt.Println("under images")
            images ++
        }
    }else if n.Type == html.TextNode {
        text := strings.TrimSpace(n.Data)
        lineScanner := bufio.NewScanner(strings.NewReader(text))
        lineScanner.Split(bufio.ScanLines)

        for lineScanner.Scan() {
            wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
            wordScanner.Split(bufio.ScanWords)

            for wordScanner.Scan() {
                words++
            }

            if err := wordScanner.Err(); err != nil {
                fmt.Println(err)
            }
        }

        if err := lineScanner.Err(); err != nil {
            fmt.Println(err)
        }
    }

    for c:=n.FirstChild; c != nil; c = c.NextSibling{
        w, i := countWordsAndImages(c)
        words += w
        images += i
    }

    return 
}
