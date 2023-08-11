package main

import (
	"fmt"
	"time"
    "flag"
)

type employee struct{
    email string
    id int
}
func f()*int {
    v := 1
    return &v
}


type newint int
func accepinterface(v any){
    fmt.Printf("%v %T\n", v, v)
    newV := v.(newint)
    fmt.Printf("%v %T\n", newV, newV)

    switch v.(newint) {
        case 


}

func main() {
    var name []string
    var ptr *int
    var m map[string]int
    var emp employee 
    fmt.Printf("%+v\n", name, )
    fmt.Printf("%v, %p\n", ptr, ptr)
    fmt.Printf("%v\n", m)
    fmt.Printf("%v %T\n", emp, emp)
    // p := &x
    // fmt.Printf("pointer %d %T\n", p, p)


    var p = f()
    fmt.Print(f() == f())
    fmt.Printf("%v, %T\n", p, p)
    fmt.Printf("%v, %T\n", *p, *p)

    var q = f()
    fmt.Printf("%v, %T\n", q, q)
    fmt.Printf("%v, %T\n", *q, *q)
    
    fmt.Println("time and command line")
    
    tenHour, _ := time.ParseDuration("10h")

    var newT time.Duration
    flag.DurationVar(&newT, "t", tenHour, "time to sleep")
    flag.Parse()
    fmt.Println("show", newT)
    time.Sleep(newT)

    
    var newx interface{} = 10
    accepinterface(newx)
}
