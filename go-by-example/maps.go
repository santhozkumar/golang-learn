package main

import "fmt"


func main() {

    d := make(map[string]int)
    d["k1"] = 1
    d["k2"] = 2


    delete(d, "k2")


    _, prs := d["k2"]
    fmt.Println(d, prs)

}
