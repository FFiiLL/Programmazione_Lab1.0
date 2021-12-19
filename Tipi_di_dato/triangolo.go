package main

import(

    "fmt"
    "math/rand"
    "math"

)
const EPSILON = 1.0e-9



func main() {

    var s int64
    
    var m1,q1,m2,q2,m3,q3 float64
    
    fmt.Print("inserisci s:")
    fmt.Scan(&s)
    fmt.Print("inserisci m1e q1:")
    fmt.Scanf("%f %f",&m1,&q1)
    fmt.Print("inserisci m2 e q2:")
    fmt.Scanf("%f %f",&m2,&q2)
    fmt.Print("inserisci m3e q3:")
    fmt.Scanf("%f %f",&m3,&q3)
    rand.Seed(s)
    for i:=0;i<10;i++{
    
        px := rand.Float64() *10.0
        py := rand.Float64() *10.0
        
        r1 := px * m1 + q1
        r2 := px * m2 + q2
        r3 := px * m3 + q3
        
        if ÈXMaggioreDiY(py, r2) || ÈXMaggioreDiY(py, r3) || ÈXMinoreDiY(py, r1) {
            fmt.Println(px,py,"esterno")
        }else{
            
            fmt.Println(px,py,"INTERNO AL TRIANGOLO")
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
