package main


import (
    "fmt"
    "unicode/utf8"
)
    


func main () {

    const sample string = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    const sample2 string = "\xe0\xae\x8a"


    fmt.Println(sample2)
    fmt.Println(sample)
    
    fmt.Printf(" %x\n", sample)
    fmt.Printf("%q\n", sample)
    fmt.Printf("%+q\n", sample)
    for i:=0;i<len(sample);i++{
        fmt.Printf(" %x\n", sample[i])
    }



    by_slice:= []byte{'\xe2', '\x8c', '\x98'}

    for i:=0; i<len(by_slice); i++ {
        fmt.Printf("% x byte\n", by_slice[i])
    }
    

    const placeOfInterest = `⌘\x20`


    fmt.Printf("Plain String : ")
    fmt.Printf("\n%s\n", placeOfInterest)



    fmt.Println("thai")
     const nihongo = "日本語"

     for index, runevalue := range nihongo {
         fmt.Printf("%#U string starts at %d\n", runevalue, index)

     }


    for i, w := 0, 0; i < len(nihongo); i+=w {
        runeval, width := utf8.DecodeRuneInString(nihongo)
        fmt.Printf("%#U string starts at %d\n", runeval, i)

        w = width

    }

    fmt.Println("count")
    fmt.Println(len(nihongo))
    fmt.Println(utf8.RuneCountInString(nihongo))
}
