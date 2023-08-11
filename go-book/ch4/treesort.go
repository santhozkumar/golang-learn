package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}


func add(t *tree, value int) *tree {

	if t == nil {
		t = new(tree)
		t.value = value
        return t
	}
    if value < t.value {
        t.left = add(t.left, value)
    }else{
        t.right = add(t.right, value)
    }

	return t
}



func Sort(values []int){
    var root *tree
    for _, value := range values{
        root = add(root, value)
    }

    appendValues(values[:0], root)
}


func appendValues(values []int, t *tree) []int{

    if t!=nil{
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }

    return values

}


func main(){

    var t *tree
    fmt.Println(t)
    t = add(t, 10)
    t = add(t, 20)
    t = add(t, 30)

    fmt.Println(t)
    fmt.Println(t.left)
    fmt.Println(t.right)
    fmt.Println(t.right.right)

    // values = make([]int, 10)
    values := []int{83,72,9,10,1,100}
    fmt.Println(values[:0])
    Sort(values)
    fmt.Println(values)
}

