package main

import (
    "fmt"
    "time"
)


const NumWorkers = 2 

type Work struct {
    x, y, z int
}

func worker(in chan *Work, out chan *Work) {
    for w := range in {
    // w := <- in
        w.z = w.x * w.y
        time.Sleep(2 * time.Second)
        fmt.Println("back from sleep")
        out <- w
    }
    close(out)
}


func sendLotsofWork(in chan *Work){
    for i := 0; i < 8; i++ {
        fmt.Println(i)
        in <- &Work{x:4, y:3}
    }
    close(in)
}

func receiveLotofWork(out chan *Work){
    for elem  := range out {
        fmt.Println("received out: ", elem)
    }
    // for i := 0; i < 5; i++ {
    //     <- out
    // }
} 

func main() {
    in, out := make(chan *Work), make(chan *Work)
    for i :=0; i <NumWorkers; i++ {
        go worker(in, out)
    }
    go sendLotsofWork(in)
    receiveLotofWork(out)

}
