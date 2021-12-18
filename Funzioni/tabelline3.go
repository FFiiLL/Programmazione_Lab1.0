package main

import(

    "fmt"

)

func Tabellina(numero int) bool {
    
    if numero >=1 && numero <=9{
        fmt.Print("Tabellina del numero ",numero,": ")
        for i :=0 ; i<=10 ; i++{
            fmt.Print(i*numero," ")
        }
        fmt.Println()
        return true
    }else{
        fmt.Println("PROGRAMMA TERMINATO")
        return false
    }
    

}


func main() {
    
    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
     for Tabellina(n){
     
        fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
     }

}
