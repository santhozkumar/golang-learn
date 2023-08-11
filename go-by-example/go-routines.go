package main

import (
	"fmt"
	"time"
)




func dummy(s string) {

    for i:=0; i<3; i++ {

        fmt.Println(s, i)
        time.Sleep(time.Second)
    }

}

// func main() {
//
//
//     fmt.Println("start")
//
//     go dummy("something")
//     go dummy("Second")
//
//
//
//     time.Sleep(time.Second)
//
//
//     go func(msg string) {
//         fmt.Println(msg)
//     }("guest")
//
//
//     time.Sleep(time.Second)
//
//     fmt.Println("done")
// }


func main() {

    messages:=make(chan string, 1)

    go func(){
        for i:=0;i<9;i++ {
            time.Sleep(time.Second)
            fmt.Print(i)
            if i == 4 {
                fmt.Println("inside the loop")
                messages <- "entered"
                messages <- "second"
            } else if i ==6 {
                fmt.Println("inside the 6 condition")
                messages <- "6 entered"
            }
        }
        }()


    msg := <- messages
    fmt.Println(msg)
    msg = <- messages
    fmt.Println(msg)
    fmt.Println("done")
}






// Channel directions 

func send_ping(ping chan<- string, message string) {
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    ping <- message
}

func receive_pong(ping <-chan string, pong chan<- string) {

    fmt.Println("receive_pong started")
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    
    fmt.Println(<-ping)

    pong <- "finished pong"


}

// func main() {
//     ping := make(chan string, 1)
//     pong := make(chan string, 1)
//
//     go send_ping(ping, "pinging here")
//
//     go receive_pong(ping, pong)
//
//     fmt.Println(<-pong)
//
//     fmt.Println("done")
//
//
//
// }





