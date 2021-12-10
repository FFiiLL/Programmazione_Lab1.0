package main

import(

    "fmt"

)

func main() {
    
    var a1, a2 int
    
    fmt.Print("inserire le ampiezze dei due angoli: ")
    fmt.Scan(&a1,&a2)
    
    if a1+a2 >= 180 || a1 <= 0 || a2 <= 0{
        fmt.Println("i due angoli non appartengono a un triangolo")
        }else{
        
            fmt.Println("ampiezza terzo angolo:",180-a1-a2)
        
        }
    
    

}
