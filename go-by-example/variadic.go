package main


import "fmt"

func concat(separator string, listofStrings ...string) string {
    var result string
    for i, s:= range(listofStrings) {
        if i >0 {
            result += separator
        }
        result += s

        }
    return result
}


func main() {
    fmt.Println("start")
    output := concat(" ", "hello", "world")
    fmt.Println(output)
}
