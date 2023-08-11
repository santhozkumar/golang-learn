package main

import (
	"fmt"

)



type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    val T 
    next *element[T]
}



func (lst *List[T]) Push (v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val:v}
        lst.tail = lst.head
    }    else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}




func (lst * List[T]) GetAll () []T {

    var elements []T
    for i:=lst.head; i != nil; i = i.next {
        elements = append(elements, i.val)
    }
    return elements
}


func main() {

    fmt.Println(10)



    lst := List[int]{}
    lst.Push(10)
    lst.Push(11)
    fmt.Println(lst)

    fmt.Println(lst.GetAll())

    }




