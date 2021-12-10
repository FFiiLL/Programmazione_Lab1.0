package main

import(
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(int64(time.Now().Nanosecond()))
    var numeroCasuale int = rand.Intn(100)
    var cont int = 1
    
    for{
    
        var n int 
        fmt.Print("tentativo n",cont,": ")
        fmt.Scan(&n)
        cont++
        
        if n == numeroCasuale{
        
            fmt.Println("Hai indovinato in",cont-1,"tentativi")
            break
        
        }else if n > numeroCasuale{
        
            fmt.Println("troppo alto riprova")
            
            }else{
                fmt.Println("troppo basso riprova")
            }   
    }
    
    
    
    
    
    
}    
