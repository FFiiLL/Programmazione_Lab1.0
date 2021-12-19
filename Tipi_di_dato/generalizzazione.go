package main

import(

    "fmt"
    "math"

)

func main() {

    var x float64 
    var n int
    
    fmt.Scan(&x)
    fmt.Scan(&n)
    
    cifra := math.Pow(10,float64(n))
    
    fmt.Print("Valore troncato ",math.Trunc(x*cifra)/cifra)
    fmt.Println()
    fmt.Print("Valore arrotondato ",math.Round(x*cifra)/cifra)
    fmt.Println()

}
