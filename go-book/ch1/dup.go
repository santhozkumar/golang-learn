// package main
//
//
// import (
//     "fmt"
//     "bufio"
//     "os"
// )
//
//
// func main() {
//     counts := make(map[string]int)
//     // input := bufio.NewScanner(os.Stdin)
//     input := bufio.NewScanner(os.Stdin)
//     
//     for input.Scan() {
//         counts[input.Text()] += 1
//     }
//
//     for line, n := range counts{
//         fmt.Println(line, n)
//     }
//     fmt.Println(counts)
//
//
// }


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
