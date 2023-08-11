package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DOB       time.Time
	Position  string
	Salary    int
	ManagerId int
}

func EmployeeById(id int) *Employee {
    var emp Employee
    return &emp
}


type Point struct{x, y int}

type Circle struct{
    Point Point
    Radius int
}




func main() {
	// fmt.Println("done")
	//
	//
 //    var dilbert Employee
 //    fmt.Println(dilbert.Salary)
 //    dilbert.Salary -= 5000
 //    fmt.Println(dilbert.Salary)
 //    
 //    dilbert.Position = "Manager"
 //    fmt.Println(dilbert.Position)
 //    position := &dilbert.Position
	//
 //    *position = "Senior " + *position
 //    fmt.Println(dilbert.Position)
	//
 //    fmt.Println(EmployeeById(0))



    // struct embedding 


    p := Point{10,5}

    c1 := Circle{Point:p, Radius:20}


    fmt.Printf("%#v", c1)




}
