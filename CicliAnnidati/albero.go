package main

import(

    "fmt"

)

func main() {

    var n int
    fmt.Scan(&n)
    //chioma
    for i:=0;i<n;i++{
        for j:=0;j<2*n-1;j++{
            if i == n-1 || j == (n-1)+i || j == (n-1)-i || (j <= (n-1)+i && j >= (n-1)-i) {
                fmt.Print("*")
            
            }else{
                fmt.Print(" ")
            }
        }
        fmt.Println()
    
    }
    
    //tronco
    var StartTronco int = ((n*2-1) - 3)/2
    for i:=0;i<3;i++{
       
        for j:=0;j<2*n-1;j++{
        
            if j == StartTronco || j == StartTronco + 1 || j == StartTronco - 1{
                fmt.Print("+")
            }else{
                fmt.Print(" ")
            }
        
        }
        fmt.Println()
    
    }
    
    
}			
