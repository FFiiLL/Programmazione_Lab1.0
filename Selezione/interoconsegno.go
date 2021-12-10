package main

import(

    
    
    "fmt"

)

func main() {

    var n int
    
    fmt.Print("Inserisci numeri: ")
    fmt.Scan(&n)
    
    if n <= 0{
        fmt.Println(n)
    }else{
        fmt.Print("+",n,"\n")
    }

}
