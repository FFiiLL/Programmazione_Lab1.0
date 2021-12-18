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

func SonoAmichevoli(n,m int) bool {

    if SommaDivisoriPropri(n)==m && SommaDivisoriPropri(m)==n{
        return true
    }else{
        return false
    }



}

func NumeriAmichevoli(limite int) {

    for i:=1 ; i<limite-2 ; i++{
      for j:=i+1; j<limite ; j++{ 
            if SonoAmichevoli(i,j){
                fmt.Print("(",i,",",j,")")
            }
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
        fmt.Print("Numeri : ")
        NumeriAmichevoli(n)
        fmt.Println()
    }

}
