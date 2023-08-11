package main


import "fmt"


func nextInt() func() int {
    i := 0
    return func() int {
        i ++
        return i
    }
}



func main() {

    nexti := intSeq2()

    //
    // fmt.Println(nexti())
    // fmt.Println(nexti())
    // fmt.Println(nexti())
    // fmt.Println(nexti())
    //
    //
    // next2 := nextInt()
    // fmt.Println(next2())
    // fmt.Println(next2())
    // fmt.Println(next2())
    //
    fmt.Println(nexti())
    fmt.Println(nexti())
    fmt.Println(nexti())

}



func intSeq2() func() int {
    i := 0 
    return func() int {
        i++
        return i
    }
}



