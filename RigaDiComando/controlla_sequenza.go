package main

import(

    "os" 
    "fmt"
    "strconv"
    //"math"
)

func main() {
    flag := true
    sl := make([]int,len(os.Args)-1)
    for i:=1;i<len(os.Args);i++{
       
       n, err := strconv.Atoi(os.Args[i])
       if err == nil{
            sl[i-1] = n
       }else{
            fmt.Print("valore posizione ",i-1," non valido\n")
            flag = false
            break
       }
       
    }
    
    
    
    if flag{
    
        for i, numero := range sl{
            
            if i%2!=0{
                if !(numero<sl[i-1]){
                    fmt.Print("valore posizione ",i," non valido\n")
                    flag = false
                    break
                }
            }else{
                
                if i > 0{
                    if !(numero>sl[i-1]){
                        fmt.Print("valore posizione ",i," non valido\n")
                        flag = false
                        break
                    
                    }
            
               }
 
             }       
       }
    
        if flag{
            fmt.Println("sequenza valdia")
        }
    
    }
}
