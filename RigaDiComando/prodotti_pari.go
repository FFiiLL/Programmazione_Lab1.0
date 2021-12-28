package main

import(

    "fmt"
    "os"
    "strconv"
)

func main() {

    numeri := LeggiNumeri()
    fmt.Println(Calcola(numeri))

}

func LeggiNumeri() (numeri []int){

    for i:=1 ; i<len(os.Args) ; i++{         
        if n, err := strconv.Atoi(os.Args[i]) ; err == nil{
             numeri= append(numeri,n)
        }           
    }
    return
}

func Calcola (sl []int) int {
    var somma int
    for i, numero := range sl{
    
        for _, numero2 := range sl[i+1:] {
            
            if (numero * numero2) % 2 == 0{
                somma += numero * numero2
                
            }
        
        }
    
    }

    return somma
}
