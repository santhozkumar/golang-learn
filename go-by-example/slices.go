package main

import (
    "fmt"
//    "reflect"
)

//
// func main () {
//
//     var s = make([]string, 3)
//
//     fmt.Println(len(s), s)
//     s[0] = "a"
//     s[1] = "b"
//     s[2] = "c"
//
//
//     c := make([]string, len(s))
//     copy(c, s)
//     c = append(c, "d")
//     fmt.Println("copy", c)
//     fmt.Println("original", s)
//
//     s = append(s, "d", "e")
//     fmt.Println(len(s), s)        
//
//
//     twoD := make([][]int, 3)
//     for i:=0; i<3; i++ {
//         inner_len := i+1
//         twoD[i] = make([]int, inner_len)
//         for j:=0; j<inner_len; j++ {
//             twoD[i][j] = i+j
//         }
//     }
//     fmt.Println(twoD)
//     twoD = append(twoD, []int{3,4,5,6})
//     fmt.Println(twoD)
//
//
//
//     fmt.Println("new ")
//     x := make([]byte, 5, 10)
//     fmt.Println(x, len(x), cap(x))
//
//
//     b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
//     fmt.Println(b)
//
//     r := [3]string{"Лайка", "Белка", "Стрелка"}
//     y := r[:]
//     fmt.Println(y)
//
//
// //    for i := range x {
//
// }

func main() {
    var s = []string{"a", "b"}
    fmt.Println(s)

    var newS = make([]string, len(s))
    copy(newS, s)
    fmt.Println(newS)

}
