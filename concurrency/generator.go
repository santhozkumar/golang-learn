package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIN (input1 <-chan string, input2 <-chan string) <-chan string{
    c := make(chan string)
    go func() { for {c <- <- input1 } }()
    go func() { for {c <- <- input2 } }()

    return c
    
}

func boring(msg string) <-chan string {


    c := make(chan string)
    go func() {
        for i:=0; ; i++{
            time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
            c <-fmt.Sprintf("%s %d", msg, i)
        }

    }()
    return c
}


func main() {
    // mac := boring("boring sent mac")
    // joe := boring("boring sent joe")
    
    c := fanIN(boring("boring sent mac"), boring("boring sent joe"))

    for i:=0; i<10; i++ {
        // fmt.Printf("%s \n", <-mac)
        // fmt.Printf("%s \n", <-joe)
        fmt.Printf("%s \n", <-c)
        
    }
}

