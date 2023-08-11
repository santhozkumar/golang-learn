package helloworld

import (
    "testing"
)

func TestHello(t *testing.T){
    t.Run("Saying hello to person", func(t *testing.T){

        got := Hello("santhosh", "")
        want := "Hello World santhosh"

        assertCorrectMessage(t, got, want)
        // if got != want {
        //     t.Errorf("got %q want %q", got, want)
        // }
    })

    t.Run("Saying hello with empty string", func(t *testing.T){

        got := Hello("", "")
        want := "Hello World"
        assertCorrectMessage(t, got, want)

        // if got != want {
        //     t.Errorf("got %q want %q", got, want)
        // }
    })
    t.Run("In spanish", func(t *testing.T){
        got := Hello("elona", "spanish")
        want := "Hola Worlde elona"
        assertCorrectMessage(t, got, want)
    })

    t.Run("just my understanding", func(t *testing.T) {


        
    })



}



func assertCorrectMessage(t testing.TB, got, want string){
    // t.Helper()
    if got != want{
        t.Errorf("got %q want %q", got, want)
    }
}

