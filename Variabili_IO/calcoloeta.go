package main

import(

    "fmt"
    "math"
)


func main() {

    var e1,e2 int
    fmt.Print("Eta persona 1 = ")
    fmt.Scan(&e1)
    fmt.Print("Eta persona 2 = ")
    fmt.Scan(&e2)
   
    sommaEta := e1 + e2
    mediaEta := float64(sommaEta)/2
    mediaDifetto := math.Floor(mediaEta)
    mediaEccesso := math.Ceil(mediaEta)
    somma10anni := sommaEta + 20
    media10anni := float64(somma10anni) / 2
    
    fmt.Println("somma eta =",sommaEta)
    fmt.Println("media eta =",mediaEta)
    fmt.Println("media difetto intera =",mediaDifetto)
    fmt.Println("media eccesso intera =",mediaEccesso)
    fmt.Println("somma tra 10 anni =",somma10anni)
    fmt.Println("media tra 10 anni =",media10anni)
}
