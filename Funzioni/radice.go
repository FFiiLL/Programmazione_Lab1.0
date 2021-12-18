package main

import(

    "fmt"
    "math"

)
func RadiceQuadrata( numero float64) (float64, bool) {

    if numero < 0 {
    
        return 0, false
    
    }else{
    
        return math.Sqrt(numero), true
    
    }


}


func main() {

        var n float64
        fmt.Print("Inserisci un numero: ")
        fmt.Scan(&n)
        radice, valid := RadiceQuadrata(n)
        
        if valid {
        
            fmt.Println("Radice quadrata:",radice)
        
        }else{
        
            fmt.Println("Il valore inserito non appartiene al dominio della funzione")
        
        }

}

