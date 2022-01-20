package main

import(
    
    "os"
    "fmt"
    

)

func main() {

    parola := os.Args[1]
    
    
    var lettere []string
    for _,lettera := range parola{
    
        lettere = append(lettere,string(lettera))
    }
    
    
    
    lunghezza := len(lettere)
    
    meta := len(lettere)/2 + 1
    
    
    for i:=lunghezza-1 ; i>=0 ;i--{
    
        if i+1 == meta{
            for k:=lunghezza-1 ; k>=0 ;k--{
                
                fmt.Print(string(lettere[k]))
            }
            fmt.Println()
        }else{
            
            for j:=0 ; j<lunghezza ;j++{
                if j+1 == meta{
                    fmt.Print(string(lettere[i]))
                }else{
                    fmt.Print(" ")
                }
                
            }
         
            fmt.Println()
            
        }
        
    }
    
}
