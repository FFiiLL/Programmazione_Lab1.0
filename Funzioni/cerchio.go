package main

import(

    "fmt" 
    "math"

)

func CalcolaArea(r float64) float64 {
    
    return r*r*math.Pi

}
func CalcolaCirc(r float64) float64 {
    
    return r*2*math.Pi

}



func main() {

    var raggio float64
    fmt.Print("inserisci il raggio del cerchio: ")
    fmt.Scan(&raggio)
    fmt.Printf("Area del cerchio: %f\nCirconferenza del cerchio: %f\n",CalcolaArea(raggio),CalcolaCirc(raggio))
    

}
