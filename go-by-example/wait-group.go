package main


import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i:=0; i<3; i++ {
        wg.Add(i)
        i:=i

        go func (){
            defer wg.Done()
            worker(i)
        }()

}
    wg.Wait()

    fmt.Println("fully done")
}





