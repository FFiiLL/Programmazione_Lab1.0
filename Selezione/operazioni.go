package main

import(
    "fmt"
    "math"
)

func main() {

     var x, y float64
     
     fmt.Println("Inserisci due numeri interi:")
     fmt.Scan(&x,&y)
     var massimo, minimo float64
     if x >= y{
        massimo = x
        minimo = y
     }else{
        massimo = y
        minimo = x
     }
     
     somma := x + y
     differenza := massimo - minimo
     quoz := x/y
     prod := x*y
     media := (x+y)/2
     var potenzaFor float64 = 1
     for i:=0 ; i<int(y) ; i++{
     
       potenzaFor*=x
     
     }
     
     potenzaMathpow := math.Pow(x,y)
     
     
     fmt.Println(massimo,minimo,somma,differenza,quoz,prod,media,potenzaFor,potenzaMathpow)
    
}
