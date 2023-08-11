package main

import (
	"fmt"
	"sort"
)



func equal(a map[string]int, b map[string]int) bool{

    if len(a) != len(b) {
        return false
    }

    for k, av:= range a{
        if bv, ok := b[k]; !ok || av!=bv {
            return false
        }
    }
    return true
}


func main() {
    employees := make(map[string]int)
    employees["santhosh"] = 10
    employees["revanth"] = 20
    employees["sathish"] = 30


    var lencheck  [4]int


    fmt.Printf("Lencheck %d, %v\n",len(lencheck), lencheck)

    fmt.Println(employees)

    for name, id := range employees {
        fmt.Println(name, id)

    }

    names := make([]string, len(employees))

    for name := range employees{
        names = append(names, name)
    }

    sort.Strings(names)
    fmt.Print(names)

    for _, name := range names {
        fmt.Println(name, employees[name])
    }

    fmt.Println(equal(map[string]int{"A": 42}, map[string]int{"A": 42}))



}
