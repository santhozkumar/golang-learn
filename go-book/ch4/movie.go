package main 

import (
    "encoding/json"
    "fmt"
    "log"
)



type Movie struct {
    Title string
    Color bool `json:"color,omitempty"`
    Year int `json:"released"`
    Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func check(e error){
    if e != nil{
        panic(e)
    }

}

func main() {
    var a []int
    a = append(a, 10, 11, 12)
    fmt.Println(a)

    mov := make([]Movie, 5)
    
    movie_enc, err := json.Marshal(movies)
    check(err) 
    fmt.Printf("%s\n", movie_enc)
    movie_indent, err := json.MarshalIndent(movies, "1", "    ")
    check(err) 
    fmt.Printf("%s\n", movie_indent)

    if err := json.Unmarshal(movie_enc, &mov); err !=nil {
        log.Fatalf("error unmarshaling %s", err)
    }


    fmt.Printf("%#v type %T\n", mov, mov)
}
