package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
    "strings"
)



func main() {
    // resp, err := http.Get("https://www.somenotworkingakdf.com")

    for _, url := range os.Args[1:]{
        if !strings.HasPrefix(url, "https://") {
            url = "https://" + url
        }
        resp, err := http.Get(url)
        if err != nil{
            fmt.Fprintf(os.Stdout, "got error %v", err)
        }
        fmt.Fprintf(os.Stdout, "response: %v\n", resp.Body)


        // data, err := io.ReadAll(resp.Body)
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil{
            // fmt.Fprintf(os.Stdout, "got error while reading the body %v", err)
            fmt.Printf("got some error")
        }
        fmt.Fprintf(os.Stdout, "body : %s\n%s", "nothing", resp.Status)
    }
}
