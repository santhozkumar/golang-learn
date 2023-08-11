package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte
func init() {
for i := range pc {
pc[i] = pc[i/2] + byte(i&1)
}
}




func main(){
    a := byte(1&1)
    b := byte(2&1)
    c := byte(3&1)
    d := byte(4&1)
    e := byte(5&1)
    fmt.Printf("%d\n%[1]b\n", a)
    fmt.Printf("%d\n%[1]b\n", b)
    fmt.Printf("%d\n%[1]b\n", c)
    fmt.Printf("%d\n%[1]b\n", d)
    fmt.Printf("%d\n%[1]b\n", e)
    fmt.Println(pc)
    fmt.Printf("%b\n\n", pc)
    fmt.Printf("%x\n\n", pc)


}
