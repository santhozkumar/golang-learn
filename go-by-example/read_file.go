package main

import (
    "os"
    "fmt"
    // "bufio"
    // "io"
)


func check(e error) {
    if e!= nil {
        panic(e)
    }
}


func main() {


    fmt.Println("start")


    x, err := os.ReadFile("hello.txt")
    fmt.Print(string(x))
    os.Stdout.Write(x)
    check(err)



    b := make([]byte, 4)

    f, err:= os.Open("hello.txt")
    check(err)

    n, err:= f.Read(b)

    fmt.Printf("%d bytes: %s\n", n, string(b[:n]))


    // fmt.Println(os.Environ())
    fmt.Println(os.Executable())
    val, e := os.LookupEnv("aldfajl")
    // val, err:= os.LookupEnv("PATH")
    fmt.Println(val, e)

    if e != false {
        fmt.Println(val, e)
    }
    fmt.Println(val)

}
