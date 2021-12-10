package main

import(

    "fmt"

)

func main() {
    
    var n int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    if n % 10 == 0{
        fmt.Println(n,"e multiplo di",10)
     }else{
     
        fmt.Println(n,"non e multiplo di",10)
     }
}
