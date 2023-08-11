package main

import (
    "fmt"
    "time"
)


func main() {

    fmt.Print("start")

    c1 := make(chan string, 1)
    c2 := make(chan string, 1)


    go func() {
        time.Sleep(1* time.Second)
        c1 <- "one"
        }()

    go func() {
        time.Sleep(2* time.Second)
        c2 <- "two"
        }()



    for i:=0; i<2; i++ {
        select{
        case c1val := <- c1:
            fmt.Println(c1val)
        case c2val := <- c2:
            fmt.Println(c2val)
        }

    }

}

