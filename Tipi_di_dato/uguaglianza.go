package main

import(

    "fmt"
    "math"

)

func main() {

    var r float64
    fmt.Scan(&r)
    radiceQuadrata := math.Sqrt(r)
    
    if r == math.Pow(radiceQuadrata,2) {
        fmt.Println(r, "uguale a", radiceQuadrata, "*", radiceQuadrata)
    }else{
        fmt.Println(r, "diverso da", radiceQuadrata, "*", radiceQuadrata)
    }
    

}
