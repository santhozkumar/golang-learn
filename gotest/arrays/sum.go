package arrays

import "fmt"


func Sum( numbers []int) int  {
    var total int
    for _, number:= range(numbers){
        total += number
    }

    return total
    
}


func SumAll(numbersToSum ...[]int) []int{
    // fmt.Printf("%T\n %d", numbersToSum, numbersToSum)
    numberOutput := make([]int, len(numbersToSum))

    fmt.Println(numbersToSum, len(numbersToSum), "sum check")
    for index, numbers := range(numbersToSum){
        numberOutput[index] =  Sum(numbers)
    }

    return numberOutput
}

func SumAllTails(numbersToSum ...[]int) []int{
    var sums []int

    for _, numbers := range(numbersToSum){
        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            sums = append(sums, Sum(numbers[1:]))
        }
    }
    return sums

}
