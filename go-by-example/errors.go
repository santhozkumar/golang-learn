package main

import (
    "errors"
    "fmt"
)


func f1(arg int) (int, error) {

    if arg == 42 {
        return -1, errors.New("cannot be 42")
    }

    return arg, nil

}



type argError struct {
    val int 
    prob string
}


func (ae *argError) Error () string {
    
    return fmt.Sprintf("%d found int error as %s", ae.val, ae.prob)

}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "cannot receive this type"}
    }

    return arg, nil
}


func main () {

    if a, e:=f2(42); e!=nil {
        fmt.Println("got an error", e)
    }else{
        fmt.Println("got resoponse", a, e)
    }

}
