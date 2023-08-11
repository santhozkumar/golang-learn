package main
import (
    "fmt"
    "time"
)


func main() {


    c1 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
        case x := <-c1:
            fmt.Println("got response", x)
        case <-time.After(2*time.Second):
            fmt.Println("No response so timeout")

        }
    }
