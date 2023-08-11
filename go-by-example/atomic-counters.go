package main

import (
    "fmt"
    "time"
    "sync"
    "sync/atomic"
)




func main() {

    var ops uint64

    var wg sync.WaitGroup

    start := time.Now()

    for i:=0; i<50; i++ {
        wg.Add(1)

        go func(){
            for c:=0; c<1000000; c++ {
                atomic.AddUint64(&ops, 1)
                // ops = ops + 1
            }
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("ops", ops)
    end := time.Now()

    fmt.Println(end)
    fmt.Println(start)
}
