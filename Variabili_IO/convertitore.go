package main

import("fmt")

func main() {

    var distanzaKm float64
    
    fmt.Print("Distanza (Km) = ")
    fmt.Scanln(&distanzaKm)
    fmt.Println("Distanza (mi) =", distanzaKm * 0.62137)



}
