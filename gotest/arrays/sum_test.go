package arrays

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sum_all_test(t *testing.T) {

    got := SumAll([]int{1,2}, []int{2,3})
    want := []int{3, 5}

    if !cmp.Equal(got, want){
        t.Errorf("expected %v but got %v",want, got)

    }

}

func sum_all_tails_test(t *testing.T) {
    // var arrayNumbers [][]int
    // arrayNumbers = append(arrayNumbers, []int{1,2}, []int{2,3}, []int{})

    arrayNumbers := [][]int{[]int{1,2}, []int{2,3}, []int{}}
    got := SumAllTails(arrayNumbers...)
    want := []int{2, 3, 0}

    if !cmp.Equal(got, want){
        t.Errorf("expected %v but got %v",want, got)
    }

}


func any_numbers_test(t *testing.T) {
    numbers := []int{1,2,3}

    got := Sum(numbers)
    want := 6 

    if got != want {
        t.Errorf("expected %d but got %d",want, got)

    }

}

func TestSum(t *testing.T)  {
    t.Run("Any  amount of  numbers", any_numbers_test )
    
}


func TestSumAll(t *testing.T)  {
    t.Run("Any  amount of  SumAll test", sum_all_test )
    
}
func TestSumAllTails(t *testing.T)  {
    t.Run("Any  amount of  SumTailsAll test", sum_all_tails_test)
}







