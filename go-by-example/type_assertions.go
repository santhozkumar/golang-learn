package main

import (
	"fmt"
)


func performOperations (args ...interface{}) {

    for _, arg:= range args {
        switch val := arg.(type) {
        case int:
            fmt.Println("int found", val)
        case float64:
            fmt.Println("float found")
        case string:
            fmt.Println("string found")
        }
    }
}



func main() {
    performOperations(10,0.5, "hello")

}
