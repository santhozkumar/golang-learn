package main


import "fmt"


func plus (a int, b int) int {
    return a + b
}


func plusPlus(a,b,c int) int {
    return a+b+c
}


func sum(nums ...int) int {
    fmt.Print(nums)
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}



func main() {

    x := plus(10,5)
    fmt.Println(x)

    y := plusPlus(1,2,3)
    fmt.Println(y)


    nums := []int{1,2,3,4,5}
    fmt.Println(nums, "nums in main")
    
    answer :=sum(nums...)

    fmt.Println(answer)

}

