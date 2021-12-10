package main

import (
    "fmt"
    "math"
)

func main() {

    var raggio float64
    
    fmt.Print("Raggio = ")
    fmt.Scanln(&raggio)
    
    fmt.Println("Circonferenza =", raggio * 2 * math.Pi)
    fmt.Println("Area =", raggio * raggio * math.Pi)




}
