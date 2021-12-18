package main

import(

    "fmt"

)

func main() {

    var str string
    fmt.Scan(&str)
    for _, carattere := range str {
        fmt.Print(string(carattere)," ")
    }

}
