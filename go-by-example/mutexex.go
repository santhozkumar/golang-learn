package main

import (
    "fmt"
    "sync"
)



type container struct {
    mu sync.Mutex
    counter map[string]int
}



func (c *container) incr(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counter[name] += 1

}

func main(){

    c:= container{
        counter: map[string]int{"a":1, "b":1},
    }

    fmt.Println(c.counter)
    var wg sync.WaitGroup

    func_increment := func(name string, n int) {
        for i:=0; i<n; i++ {
            c.incr(name)
        }
    wg.Done()
    }
        
    wg.Add(3)
    go func_increment("a", 100)
    go func_increment("b", 100)
    go func_increment("b", 100)
    wg.Wait()

    fmt.Println(c.counter)

}

