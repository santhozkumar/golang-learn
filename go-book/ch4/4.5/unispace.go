package unispace 

import (
    "unicode"
)

func RemoveUnicodeSpace(strs []byte){
    i := 0 

    for _, s := range strs{
        spacebool := unicode.IsSpace(rune(s))
        
        if !spacebool{
            strs[i] = s
            i++
        }
    }

}
//
// func main() {
//     strs := []byte("hello world what are you doing")
//     fmt.Printf(" %s\n",strs) 
//     RemoveUnicodeSpace(strs)
//     fmt.Printf(" %s\n",strs) 
// }
