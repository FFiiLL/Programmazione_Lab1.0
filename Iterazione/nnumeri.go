package main

import(

    "fmt"
    "math"
)

func main(){
    
    var n int
    
    fmt.Scan(&n)
    fmt.Println("inserisci",n,"numeri")
    
    var somma int
    var massimo = math.MinInt64
    var minimo = math.MaxInt64
    
    var contPos, contNeg, contNull int
    
    
    for i:=0 ; i<n ; i++{
        var numero int
        fmt.Scan(&numero)
        
        somma += numero
        
        if numero < minimo{
            minimo = numero
         }
        if numero > massimo{
            massimo = numero
         }
         
         
        if numero == 0{
            contNull++
        }else if numero > 0{
            contPos++
        }else{ 
            contNeg++
         }
    
    }
    
    fmt.Println(somma,minimo,massimo,contPos,contNeg,contNull)


}
