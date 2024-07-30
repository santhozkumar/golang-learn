package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/knadh/knphone"
)

func main() {
	k := knphone.New()
	fmt.Println(k.Encode("ಅಂಕೆಸಂಖ್ಯೆ"))
	fmt.Println(k.Encode("ಊಱಿಸಾಱು"))
	httputil.DumpRequest()
}
