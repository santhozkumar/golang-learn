package main

import (
    "fmt"
    "os"
    "regexp"
)


func append_string (slice []string, data ...string) []string {

    m := len(slice)
    n := m+len(data)

    if n > cap(slice) {
        t := make([]string, cap(slice) * 2)
        copy(t, slice)
        slice = t
    }


    copy(slice[m:n], data)
    return slice

}


func main () {
    x := []string{"hello", "world"}
    x = append_string(x, "santhosh")

    fmt.Println(x)

    x = append_string(x, []string{"what", "are", "you", "doing"}...)

    fmt.Println(x)


    filebytes, _ := os.ReadFile("hello.txt")

    
    var digitRegexp = regexp.MustCompile("[0-9]+")

    
    fmt.Println(filebytes)
    findslice := digitRegexp.Find(filebytes)
    findsl := append(make([]byte, 0), findslice[:]...) 
    fmt.Println(findsl)

}

