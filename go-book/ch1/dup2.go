package main



import (
	"bufio"
	"fmt"
	"os"
)


func countLines(file *os.File, counts map[string]int) {
    input := bufio.NewScanner(file)

    for input.Scan() {

        counts[input.Text()] ++
    }

}



func main() {
    counts := make(map[string]int)
    fmt.Println(counts)

    files := os.Args[1:]
    if len(files) == 0{

        countLines(os.Stdin, counts)
    }

    for _,arg := range files {
        f, err := os.Open(arg)
        if err != nil {
            fmt.Println("printing error")
            fmt.Println(err, f)
            continue
            }
        countLines(f, counts)
        f.Close()
    }

    for line,n := range counts {
        if n>1 {
            fmt.Println("line and count ", line, n)
        }
    }


}

