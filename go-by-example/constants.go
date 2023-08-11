package main

import (
    "fmt"
    "math"
)

var s string = "constant"

func main() {
    fmt.Println(s)
    const n = 5000
    
    fmt.Println(s)
    const d = 3e20/n

    fmt.Println("...........")
    fmt.Println(d)
    fmt.Println(int64(d))


    fmt.Println(math.Sin(n))

    i:=1
    for i < 3 {
        fmt.Println(i)
        i = i + 1
    }


    for j:=1; j<5; j++ {
        fmt.Println(j)
    }


    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n<=5; n++ {
        if n%2==0 {
            fmt.Println(n)
        }
    }

}
