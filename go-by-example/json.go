package main

import (
	"encoding/json"
	"fmt"
)


type person struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Email string `json:"liame"`
}


func main() {

    bolB, _ := json.Marshal(true)
    strB, _ := json.Marshal("hello world")
    intB, _ := json.Marshal(10)
    floatB, _ := json.Marshal(2.85)

    empty_string := []byte("")

    slice_string := []string{"hello", "world"}
    sliceB, _ := json.Marshal(slice_string)
 
    mapD := map[string]int{"hello" : 1, "world": 2}
    mapDB, _ := json.Marshal(mapD)


    fmt.Println("Map : ", string(mapDB))
    fmt.Println("Map : ", mapDB)

    fmt.Println(sliceB)
    fmt.Println(empty_string)
    fmt.Println(bolB)
    fmt.Println(strB)
    fmt.Println(intB)
    fmt.Println(floatB)
    fmt.Println(bolB)
    fmt.Println(json.Marshal(1))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    
    fmt.Println("Map : ", string(byt))
    fmt.Println("Map : ", byt)




    // var dat map[string]interface{}



    fmt.Println()




    santhosh := person{Name: "santhosh", Age: 26, Email: "santhosh@gmail"}
    santhosh_json, _ := json.Marshal(santhosh)

    fmt.Println(string(santhosh_json))
}
