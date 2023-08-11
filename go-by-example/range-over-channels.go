package main


import "fmt"

func main() {
    queue := make(chan string, 2)

    queue <- "hello"
    queue <- "world"
    
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
        
    }
}
