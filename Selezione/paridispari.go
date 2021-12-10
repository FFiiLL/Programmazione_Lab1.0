package main

import(

    "fmt"

)

func main() {
    var n int
    fmt.Print("inserisci un numero: ")
    fmt.Scan(&n)
    
    if n % 2 == 0{
        fmt.Println(n,"pari")
        }else{fmt.Println(n,"dispari")}
}
