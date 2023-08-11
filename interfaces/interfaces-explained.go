package main


import (
    "fmt"
    "strconv"
)


type Book struct {
    Name string
    author string
    cityid int
}



func (b Book) String() string {
    return b.Name + b.author + strconv.Itoa(b.cityid)
}


type new int 


func (r new) String() string {
    return strconv.Itoa(int(r))
}


func accept(x fmt.Stringer){

    
    fmt.Println(x.String())
}

func main() {
    var x new
    x = 5
    fmt.Println(x.String())

    s := "not a number"

    if _,err := strconv.ParseFloat(s, 64); err!=nil {
        fmt.Print(err, "\n", "\n")
        e := err.(*strconv.NumError)
        fmt.Println(e.Err, e.Num, e.Func)
    }


    bc := Book{Name: "ddai", author: "Martin", cityid: 10}

    // fmt.Println(bc.String())
    accept(bc)



}

