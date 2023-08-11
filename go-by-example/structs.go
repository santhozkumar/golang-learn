package main


import (
    "fmt"
)



type person struct {
    name string
    age int
    mail string
}


type rect struct {
    height, width int
}


func (r *rect)area () int {
    return r.width * r.height
}


func (r *rect) perim () int {
    return 2*r.width + 2*r.height
}



func newPerson (name string) *person {
    p := person{name: name }
    p.age = 42

    return &p
}


func main() {
    fmt.Println(person{name: "santhosh", age: 20})

    personpointer := newPerson("revenath")

    fmt.Println(personpointer.name, personpointer.age)


    r:=rect{10,5}

    fmt.Println(r.area())
    fmt.Println(r.perim())




}

