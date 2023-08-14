package shape

import "math"

type Shape interface{
    Area() float64
    Perimeter() float64
}


type Triangle struct {
    Base float64
    Height float64
}

type Rectangle struct {
    Width float64
    Height float64
}

type Circle struct {
    Radius float64
}



func (r Triangle) Area() float64{
    // return 0.5 * r.Height * r.Base
    return 0
}


func (r Triangle) Perimeter() float64{
    return 2

}
func (r Rectangle) Area() float64{
    return r.Height * r.Width
    // return 0
}


func (r Rectangle) Perimeter() float64{
    return 2 * (r.Height + r.Width)

}


func (r Circle) Area() float64{
    return math.Pi * math.Pow(r.Radius, 2)
}


func (r Circle) Perimeter() float64{
    return 2 * math.Pi * r.Radius

}

