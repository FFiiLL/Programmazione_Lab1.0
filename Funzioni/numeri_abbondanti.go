package main

import(

    "fmt"

)

func SommaDivisoriPropri(n int) int {
    var somma int
    for i:=1 ; i<n ; i++ {
        if n%i == 0 {
            somma +=i
        }
    }
    return somma
}

func EAbbondante(n int) bool{
    
    if n < SommaDivisoriPropri(n){
        return true
    }else{
        return false
    }


}

func NumeriAbbondanti(limite int) {
    
    for i:=1 ; i<limite ; i++{
        if EAbbondante(i){
            fmt.Print(i," ")
        }
    }


}

func main() {

    var n int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    if n <= 0 {
        fmt.Println("soglia non positivo")
    }else{
        fmt.Print("Numeri abbondanti: ")
        NumeriAbbondanti(n)
        fmt.Println()
    }

}
