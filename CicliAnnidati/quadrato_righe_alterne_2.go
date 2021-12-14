package main

import(

   "fmt" 

)

func main() {

    var n,flagInt int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    flagInt = 0
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            
            if flagInt == 0{
                fmt.Print("* ") 
               
            }else if flagInt == 1{
                fmt.Print("+ ")
               
            }else if flagInt == 2{
                fmt.Print("o ")
              
            }
            
            
        }
        
        if flagInt == 0{
               
                flagInt++
            }else if flagInt == 1{
               
                flagInt++
            }else if flagInt == 2{
              
                flagInt = 0
            }
        
        
        
        
        fmt.Println()
    }  

}
