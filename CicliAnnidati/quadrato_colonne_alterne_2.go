package main

import(

    "fmt"

)

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    
    for i:=0;i<n;i++{
       //caso n pari
        if n%2==0{
            for j:=0;j<n/2;j++{
                if j%2==0{
                    fmt.Print("* * ")
                }else{
                fmt.Print("+ + ")
                }
            
            }
        }else{
        
            
        for j:=0;j<(n/2)+1;j++{
                
                if j == (n/2){
                
                if j%2==0{
                    fmt.Print("* ")
                }else{
                fmt.Print("+ ")
                }
                
                
                }else if j%2==0{
                    fmt.Print("* * ")
                }else{
                fmt.Print("+ + ")
                }
            
            }
        
        
        
        
        
        
        
        
        
        
        
        }
        fmt.Println()
     } 
        
        
        
        
}
    
    
    
    
    
    


