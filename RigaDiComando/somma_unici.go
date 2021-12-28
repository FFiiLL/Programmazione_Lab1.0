package main

import(

    "os"
    "fmt"
    "strconv"

)

func LeggiNumeri() (numeri []int){

    for i:=1 ; i<len(os.Args) ; i++{         
        if n, err := strconv.Atoi(os.Args[i]) ; err == nil && n>=0 && n<=100{
             numeri= append(numeri,n)
        }           
    }
    return
}

func Occorrenze(numeri []int, n int) int{

    var cont int
    for _, numero := range numeri{
        
        if numero == n {
            cont++
        }
    
    }

    return cont
}

func main() {

    numeri := LeggiNumeri()
    var somma int
    
    for _, numero := range numeri{
        
        
        if Occorrenze(numeri,numero) == 1{
            somma += numero
        }
        
        
        
    }
    fmt.Println(somma)

}
