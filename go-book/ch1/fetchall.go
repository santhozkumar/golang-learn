package main


import (
    "fmt"
    "io"
    "time"
    "net/http"
    "os"
)



func fetch(url string, ch chan <- string){
    start := time.Now()
    resp, err := http.Get(url)
    if err!= nil{
        fmt.Fprintf(os.Stdout, "Error in get %v", err)
    }

    nbytes, err := io.Copy(io.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <-fmt.Sprintf("Got error in body %v \n\nyeah printing this \n\n", err)
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("secs : %.2f, byes: %d, url: %s", secs, nbytes, url)


}


func main(){
    start_time := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:]{
        go fetch(url, ch)
    }
    for range os.Args[1:]{
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed time\n", time.Since(start_time).Seconds())
}


