package unique



func RemoveAdjacent(strs []string) []string{
    i := 0
    for _, s := range strs{
        if strs[i] == s {
            continue
        }
        i++
        strs[i] = s
    }
    return strs[:i+1]

}


// func main() {
//     s := []string{"a", "a", "b", "c", "e", "e", "d", "d", "d"}
//     fmt.Println(s)
//     s = RemoveAdjacent(s)
//     fmt.Println(s)
//
//
// }
//
//
//
