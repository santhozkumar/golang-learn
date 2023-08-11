package main

import (
	"crypto/sha256"
	"fmt"
)

func zero(ptr *[32]byte) {
    for i := range ptr{
        fmt.Println(ptr[i], i)
        ptr[i] = 0
    }
}

func main(){

    var a [3]int
    // var b [4]int

    for i, e := range a{
        fmt.Printf("%d, %d\n", i, e)
    }

    c:= [...]int{1,2,3,4,5}
    fmt.Printf("%d, %[1]T\n", c)

    type Currency int
    const (
        USD Currency=1 << iota
        EUR
        GBP
        RMB
    )
    fmt.Println(USD, EUR, GBP, RMB)
    fmt.Print()


    s :=[...]string{0:"hello", 1:"world"}
    fmt.Printf("%#v", s)


    symbol := [...]string{USD: "dollar", EUR:"euro"}
    fmt.Println(symbol[USD], USD)

    const x = 200
    r:= [...]int{x:-1}
    fmt.Println(r)


    ru := []byte{'x'}
    ra := []byte{'X'}
    fmt.Printf("%v, %[1]b, %[1]x\n", ru)
    fmt.Printf("%v, %[1]b, %[1]X\n", ra)
    
    p_check := make([]int, 5)

    fmt.Printf("\n%p", p_check)


    abyte:= sha256.Sum256(ru)
    bbyte:= sha256.Sum256(ru)

    fmt.Printf("\n%x\n%x\n%t\n",abyte, bbyte, abyte==bbyte)


    fmt.Printf("%T",abyte) 
    zero(&abyte)

    fmt.Println(abyte)

    fmt.Printf("\n%x\n", abyte)
    fmt.Printf("\n%x\n", bbyte)

    for i := range bbyte {
        fmt.Println(i, bbyte[i])
    }

    s_test := []string{1:"jan", 2:"Feb"}
    fmt.Printf("%T", s_test)
    fmt.Println(s_test[0:3])
}
