package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
    fmt.Println(time.Now())
    time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
}
