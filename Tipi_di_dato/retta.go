package main

import(
    "fmt"
    "math/rand"
    "math"
)

const EPSILON = 1.0e-9


func main() {
    
    var s int64
    var m,q float64
    
    fmt.Print("inserisci s")
    fmt.Scan(&s)
    fmt.Print("inserisci m e q")
    fmt.Scanf("%f %f",&m,&q)
    rand.Seed(s)
  for i:=0;i<10;i++{
    
    x := rand.Float64() *10
    y := rand.Float64() *10
    
    if ÈXMaggioreDiY(y, m * x + q){
        fmt.Println(x,y,"SOPRA")
    }else if ÈXMinoreDiY(y, m * x + q){
        fmt.Println(x,y,"SOTTO")
    }else{
        fmt.Println(x,y,"FA PARTE")
    }
  
  }

   
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore 
di `y` di almeno una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
    return (x - y) > EPSILON 
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale 
a `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
    return math.Abs(x - y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore 
di `y` di almeno una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
    return (x - y) < -EPSILON
}

