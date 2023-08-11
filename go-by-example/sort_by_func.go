package main

import (
    "fmt"
    "sort"
)




type bylength []string


func (r bylength) Len() int {
    return len(r)
}

func (r bylength) Swap(i, j int) {
    r[i], r[j] = r[j], r[i]
}


func (r bylength) Less(i,j int) bool{
    if len(r[i]) < len(r[j]) {
        return true
    }
    return false

}



func main() {
    fmt.Println("start")
    st := []string{"hello", "apple", "aldkjfasd"}

    sort.Sort(bylength(st))
    fmt.Println(st)





}
