package unispace

import (
	"bytes"
	"testing"
)



func TestRemoveUnicodeSpace(t *testing.T){
    strs := []byte("hello world what are you doing") 
    RemoveUnicodeSpace(strs)
    want := []byte("helloworldwhatareyoudoingdoing")

    if !bytes.Equal(strs, want){

        t.Errorf("got %s but want %s",strs, want)
    }

}
