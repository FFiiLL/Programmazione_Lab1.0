package main

import(

    "fmt"
    "github.com/FFiiLL/Programmazione_Lab1.0/tree/main/Packages/triangolo"


)

func main() {

    
    
    fmt.Println("lati:",triangolo.NuovoTriangolo(5,4,3))
    triangolo.AreaPeri(triangolo.NuovoTriangolo(5,4,3))

}
