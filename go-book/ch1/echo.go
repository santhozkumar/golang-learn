package main

import (
    "os"
    "fmt"
    "time"
    "strings"
    )



func main() {

    start := time.Now()

    other_way := strings.Join(os.Args[1:len(os.Args)], " ")

    secs := time.Since(start).Seconds()
    fmt.Printf("%.8fs %s", secs, other_way)

    start = time.Now()
    
    var s, sep string
    for i:=1; i<len(os.Args); i++ {
        s += sep+os.Args[i]
        sep = " "
    }
    fmt.Println(s)

    secs = time.Since(start).Seconds()

    fmt.Printf("%.7fs", secs)
    
}


