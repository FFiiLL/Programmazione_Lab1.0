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

func EPerfetto(n int) bool{
    if n == 0{
        return false
    }
    if n == SommaDivisoriPropri(n){
        return true
    }else{
        return false
    }


}

func main() {

    var n int
    fmt.Print("inserisci un numero: ")
    fmt.Scan(&n)
    if EPerfetto(n) {
        fmt.Println(n,"e perfetto")
    }else{
        fmt.Println(n,"non e perfetto")
    }
}
