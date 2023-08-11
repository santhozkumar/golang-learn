package main

import (
	"fmt"

)



func mayPanic () {
    panic("error returned")

}


func main() {

    defer func(){
        if r:= recover(); r!=nil {
            fmt.Println("rever from defer")
        }

        }()


    mayPanic()
    fmt.Println("after panic")
}

