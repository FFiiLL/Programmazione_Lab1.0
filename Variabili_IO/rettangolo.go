package main 

import "fmt"

func main() {
    
    var base, altezza int
    
    fmt.Print("Inserisci la base: ")
    fmt.Scanln(&base)
    fmt.Print("Inserisci l'altezza: ")
    fmt.Scanln(&altezza)
    fmt.Println("Perimetro =", (base+altezza)*2)
    fmt.Println("Area =", base*altezza)    


}
