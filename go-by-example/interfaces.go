package main

import (
    "fmt"
    "math"
)



type geometry interface {
    area() float64
    perim() float64

}



type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}


func (r *rect) area () float64 {
    return r.width * r.height
}
func (r circle) area () float64 {
    return math.Pi * math.Pow(r.radius, 2)
}


func (r *rect) perim () float64 {
    return 2*(r.width + r.height) 
}
func (r *circle) perim () float64 {
    return 2 * math.Pi * r.radius 
}



func measure (g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())

}
func main () {
    fmt.Println()
    r := &rect{10,5}
    c := &circle{5}

    measure(r)
    measure(c)
}

