package main

import (
	// "container/heap"
	"fmt"

)

type Item struct {
    value string
    index int
    priority int
}

type priorityQueue []*Item


func (pq priorityQueue) Len() int {return len(pq)}

func (pq priorityQueue) swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq priorityQueue) Less(i, j int) bool {    
    return pq[i].priority > pq[j].priority
}
    
func (pq *priorityQueue) Push(x any) {
}

func (pq *priorityQueue) Pop()any{


    return 10

}

func main() {
    items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
        
    pq := make(priorityQueue, len(items))
    i := 0
    for value, priority := range items {
        pq[i] = &Item{
            value: value,
            index: i,
            priority: priority,
        }
        i++
    }

    // heap.Init(&pq)
    x := &Item{
        value:"hello",
        index: 10,
        priority: 1,
    }
    fmt.Printf("%T", x)
    fmt.Println(x.(*Item))


}

