package main

import(

    "fmt"
    "unicode"
)
func inMinuscolo(c rune) rune {
    return unicode.ToLower(c)
}

func inMaiuscolo(c rune) rune {
   return unicode.ToUpper(c) 
}

func main() {

    var stringa string
    fmt.Scan(&stringa)
    fmt.Print("TESTO MAIUSCOLO: ")
    for _, carattere := range stringa{
        carattere = inMaiuscolo(carattere)
        fmt.Print(string(carattere))
    }
    fmt.Println()
    fmt.Print("testo minuscolo: ")
    for _, carattere := range stringa{
        carattere = inMinuscolo(carattere)
        fmt.Print(string(carattere))
    }
    fmt.Println()
}

