package main

import "fmt"



func main() {

    jobs := make(chan int, 3)
    done := make(chan bool, 1)



    go func() {
        for {
            job, more := <-jobs

            if more {
                fmt.Println("processing ", job)
            } else {
                done <- true 
                return 
            }
        }
        }()

    for i:=0; i<3; i++ {
        jobs <- i
    }
    close(jobs)

    <- done
}



