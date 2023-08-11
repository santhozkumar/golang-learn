package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)


func main(){

    w := flag.Int("width", 256, "hash width algorith")

    flag.Parse()


    x := []byte{'x'}
    fmt.Printf("%T, %v, %[1]b, %[1]x", x,x)
    switch *w {
        case 384:
            fmt.Printf("\n%x\n%d\n", sha512.Sum384(x),  sha512.Sum384(x))
            fmt.Println("got 384")
        case 512:
            fmt.Printf("\n%x\n", sha512.Sum512(x))
            fmt.Println("got 512")
        case 256:
            fmt.Printf("\n%x\n", sha256.Sum256(x))
            fmt.Println("got 256")
        default:
            fmt.Println("incorrect")
        }
}

