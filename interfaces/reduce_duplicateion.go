package main

import (
	"encoding/json"
	"fmt"
	"io"

	// "bytes"
	"os"
)


type Name struct {
    FirstName string
    LastName string
}


func (n *Name) WriteJson(w io.Writer) {
     name_b, _ := json.Marshal(n)

    fmt.Println(name_b)

    _, err := w.Write(name_b)
    if err!=nil {
        panic(err)
    }

    fmt.Println("done")
}
    



func write_to_something (x io.Writer) {

    name := Name{"santhosh", "kumar"}

    name_b, _ := json.Marshal(name)

    fmt.Println(name_b)

    _, err := x.Write(name_b)
    if err!=nil {
        panic(err)
    }

    fmt.Println("done")
}



func main() {

    f, err := os.Create("hello.txt")
    if err!= nil {
        panic(err)
    }

    write_to_something(f)


}
