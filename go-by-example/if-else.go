package main


import "fmt"

func main() {
    
    if n := 9; n%2 == 0 {
        fmt.Printf("%d is even\n", n)
    } else if n < 10 {
        fmt.Printf("%d is less than 10\n", n)
    } else {
        fmt.Printf("%d is odd\n", n)
    }
}
