package main


import "fmt"

func main() {


    s := []int{1,2,3}
    fmt.Println(s)
    sum := 0

    for _, value := range s {
        sum = sum + value

    }


    kvs := map[string]string{"a":"1", "b":"2"}

    for k, value := range kvs {
        fmt.Print(k,value)

    }

    for i, c := range "go" {
        fmt.Println(i,c)
    }

}

