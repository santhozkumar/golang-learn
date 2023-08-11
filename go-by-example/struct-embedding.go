package main

import (
	"fmt"

)



type base struct {
    val string
}



func (v base) PrintNew () {
    fmt.Println(v.val)
}



type container struct {
    base
    str string 
}



func main () {
    x := base{val:"hello"}
    c := container{base: x, str: "new"}

    fmt.Println(c.val)
    fmt.Println(c.str)
    c.PrintNew()


    type printer interface {
        PrintNew() 
    }

    var y printer = c
    y.PrintNew()


}
