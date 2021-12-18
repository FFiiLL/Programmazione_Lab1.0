package main

import(

    "fmt"

)

 func EPrimo(n int) bool {

    flag := true
    
    for i:=2;i<n/2;i++{
        
        if n % i == 0{
            flag = false
            return flag
        }
    
    
    }

    return flag

}
func NumeriPrimiGemelli(limite int) {

    for i := 3 ; i<limite ; i++{
    
        if EPrimo(i) && EPrimo(i+2) && i+2 < limite{
            fmt.Print("(",i,",",i+2,")")
        }
    }

    fmt.Println()

}



func main() {

    NumeriPrimiGemelli(72)

}
