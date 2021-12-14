package main

import(

    "fmt"

)

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    flag := true
    
    for i:=0;i<n;i++{
       
        
        for j:=0;j<n;j++{
             if j % 2 == 0 {
            flag = false
        }else{
            flag = true
        }
            if flag{
                fmt.Print("+ ")
            }else{
                fmt.Print("* ")
            }
        
        }
    
    
    
        fmt.Println()
    
    }
    
    
    
    
    
    

}
