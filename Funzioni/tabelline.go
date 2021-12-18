package main

import(

    "fmt"

)

func Tabellina(numero int) {
    
    if numero >=1 && numero <=9{
        fmt.Print("Tabellina del numero ",numero,": ")
        for i :=0 ; i<=10 ; i++{
            fmt.Print(i*numero," ")
        }
        fmt.Println()
    }
    

}


func main() {
    
    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    Tabellina(n)

}
