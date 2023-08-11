package main
import (
    "fmt"
    "time"
)


func main() {



    requests := make(chan int, 5)

    for i:=0;i<5;i++ {
        requests <- i
    }
    close(requests)
    

    limiter := time.NewTicker(200 * time.Millisecond)
    for req := range(requests) {
        <- limiter.C
        fmt.Println("request", req, time.Now())
    }




    burstLimiter := make(chan time.Time, 3)

    for i:=0; i<3; i++ {
        burstLimiter <- time.Now()
    }


    go func () {
        for t:= range time.Tick(200*time.Millisecond){
            burstLimiter <- t
        }
    }()


    burstyRequests := make(chan int, 5)
    for i:=0;i<5;i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req:= range(burstyRequests) {
        req_time := <-burstLimiter
        fmt.Println("bursty rquest", req_time, req)
    }

}

