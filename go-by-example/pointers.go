package main


import "fmt"


func zeroptr(i *int) {
    *i = 10
}



func zeroval(i int) {
    i = 5
}



func main() {


    var a int = 0
    fmt.Println(&a)
    
    zeroval(a)

    fmt.Println(a)

    zeroptr(&a)
    fmt.Println(a)
}

