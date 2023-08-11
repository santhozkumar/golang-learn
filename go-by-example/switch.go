package main


import (
    "fmt"
    "time"
)


type myWeekDay int


// const (
//     Sunday myWeekDay = 1
//     Monday myWeekDay = 2 
//     Tuesday myWeekDay = 3
//     Wednesday myWeekDay = 4
//     Thursday myWeekDay = 5
//     Friday myWeekDay = 6
//     Saturday myWeekDay = 7
// )
const (
	Sunday myWeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
    i := 2

    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    }
    fmt.Printf("%T %v %d\n", time.Saturday, time.Saturday, time.Saturday)
    
    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
        // case Saturday, Sunday:
            fmt.Println("its a weekend")
        default:
            fmt.Println("its a weekday")
        }
    }
