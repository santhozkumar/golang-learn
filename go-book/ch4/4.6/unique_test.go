package unique

import (
    "testing"
)

func TestRemoveAdjacent(t *testing.T){
    want := []string{"a", "b", "c", "e", "d"}
    got := RemoveAdjacent( []string{"a", "a", "b", "c", "e", "e", "d", "d", "d"})

    for i := range want {
        if want[i] != got[i] {
            t.Errorf("%v incorrect %v", got, want)
        }
}
}
