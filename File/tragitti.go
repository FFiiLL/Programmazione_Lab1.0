package main

import(

   "fmt"
   "os"
   "bufio" 

)

type Punto struct {
	etichetta string
	x, y      float64
}

type Tratta struct {
	p1, p2 Punto
}

type TragittoPerTratte struct {
	id        int           // l'identificativo del tragitto
	tratte    []Tratta      // l'insieme delle tratte che definiscono il tragitto
	lunghezza float64       // la lunghezza del tragitto
}

type TragittoPerPunti struct {
	id        int           // l'identificativo del tragitto
	punti     []Punto       // la sequenza di punti che definisce il tragitto
	lunghezza float64       // la lunghezza del tragitto
}

func main() {

    

}

