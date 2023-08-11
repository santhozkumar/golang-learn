package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4
	got := Add(2, 2)
	if want != got {
		t.Errorf("wanted '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	//Output: 10
}
