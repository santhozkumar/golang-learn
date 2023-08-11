package helloworld

import (
	"fmt"

)

const (
    defaultPrefix = "Hello World"
    englishPrefix = "Hello World"
    spanishPrefix = "Hola Worlde"
    frenchPrefix = "Bonjour,"

)


func gettingprefix(language string)(prefix string){
    switch language {
    case "french":
        prefix = frenchPrefix
    case "spanish":
        prefix = spanishPrefix
    default:
        prefix = englishPrefix
    }
    return 

}

func Hello(name string, language string) string {
    if name == ""{
        return defaultPrefix 
    }
    if language == "spanish"{
        return fmt.Sprintf("%s %s", spanishPrefix, name)
    }
    return fmt.Sprintf("%s %s", englishPrefix, name)

}
