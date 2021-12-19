

package main

import(

    "fmt"
    "math/rand"
    "math"

)
EPSILON = 1.0e-6
func main() {

    var s int64
    var xC, yC float64
    var raggio float64
    fmt.Print("Inserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): ")
	fmt.Scan(&xC, &yC)

	fmt.Print("Inserisci il valore del raggio: ")
	fmt.Scan(&raggio)

	fmt.Println()
    for i:=0;i<1000000;i++{
        
        x := rand.Float64() *10
        y := rand.Float64() *10
        
        
        
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
