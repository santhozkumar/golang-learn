package main 

import ("fmt"
        hello "github.com/santhozkumar/gotest/helloworld"
    )

func main() {
    fmt.Println("something")
    x := hello.Hello("santhosh", "spanish")
    fmt.Println(x)
}

