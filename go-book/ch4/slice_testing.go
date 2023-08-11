package main

import "fmt"

func reverse[T any](s []T){

    // for i, j := 0, len(s)-1;i<len(s)-1; i, j = i+1,j-1 {
    //     s[i], s[j] = s[j], s[i]
    // }

    start := 0
    end := len(s) -1 
    for start <= end{
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}



func rotate[T any](s []T, n int, left bool) {

    switch left{
    case true:
        reverse(s)
        reverse(s[:n])
        reverse(s[n:])
    case false:
        reverse(s)
        reverse(s[:n])
        reverse(s[n:])
}
}

func rotateSingle[T any](s []T, n int, left bool)  {
    
    if n >len(s) || n <=0 {
        return 
    }

    switch left{
    case true:
        tail := make([]T, n)
        fmt.Printf("%T\n",tail)
        fmt.Printf("%T\n", s[:n])
        n := copy(tail, s[:n])
        fmt.Println(n, "number of ")
        fmt.Println(tail, s)

        copy(s[:len(s)-n], s[n:])
        fmt.Println(tail, s)
        copy(s[len(s)-n:], tail)

        fmt.Println(tail, s)
        // s = append(s, s[:n]...)
        // s = s[n:]
        // return s
    case false:
        // return s

    }
    // return s
}
func AppendInt(s []int, x int) []int {

    var z []int
    newlength := len(s) + 1
    if newlength < cap(s){
        fmt.Println("yes within")
        z = s[:newlength]
    } else {
        fmt.Println(newlength)
        z = make([]int,newlength, 2*newlength)
        fmt.Println(cap(z))
        copy(z, s)
    }

    z[newlength-1] = x

    return z
}


func remove[T any](s []T, i int) []T {
    copy(s[i:], s[i+1:])
    return s[:len(s)-1]
}


func main(){
    x := []int{1,2,3,4,5}
    string_slice := []string{"hello", "world"}
    reverse(x)
    reverse(string_slice)
    fmt.Println(x)
    fmt.Println(string_slice)

    rotate(x, 2, false)
    fmt.Println(x)
    rotate(x, 2, true)

    fmt.Println(x)

    y := []rune("Hello, BF")
    fmt.Println(y)

    inta := []int{1,2,3,4,5}
    inta = AppendInt(inta, 6)
    inta =AppendInt(inta, 7)
    inta =AppendInt(inta, 8)
    inta =AppendInt(inta, 9)

    fmt.Println(inta)
    
    
    x = []int{1,2,3,4,5}
    // newx := remove(x, 2)
    // fmt.Println(x) 
    // fmt.Println(newx) 

    fmt.Println("hellow")
    rotateSingle(x, 2, true)
    fmt.Println(x) 

}
