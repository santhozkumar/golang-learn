package main

import (
	"fmt"

)



func reverse_array[T any](input []T) []T{
    l := len(input)
    r := make([]T, l)

    for i, val := range input {
        r[l-1-i] = val 
    }

    fmt.Println(r)

    return r

}



func main() {

    reverse_array([]int{1,2,3,4,5})
    reverse_array([]string{"hello", "world", "santhosh"})
}

