package main

import (
    "fmt"
    "time"
)



func main() {
    
    ticker := time.NewTicker(1000*time.Millisecond)
    done := make(chan bool, 1)




    go func(){
        for {
            select {
                case <-done:
                    return
                case t:= <-ticker.C:
                    fmt.Println(t)


                }
            }
        }()


    time.Sleep(time.Second*10)
    ticker.Stop()
    done <- true



    }
