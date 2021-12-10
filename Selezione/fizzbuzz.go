package main

import (
    "fmt"
)

func main() {

    var n int
    
    fmt.Scan(&n)
    
    for i:=1 ; i<=n ; i++{
        if i % 3 != 0 && i % 5 != 0{
        
            fmt.Print(i," ")
            continue
        }else {if i % 3 == 0{
         if i % 5 != 0{   
            fmt.Print("Fizz ")
           }else{
           fmt.Print("Fizz")
           } 
        }
        if i % 5 == 0{
            fmt.Print("Buzz ")
            
        }
        }
    }



}
