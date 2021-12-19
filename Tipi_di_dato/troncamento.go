package main

import(

    "fmt"    
    "math"
)

func main() {

    var x float64 
    
    fmt.Scan(&x)
    fmt.Print("Valore troncato ",math.Trunc(x*100)/100)
    fmt.Println()

}
