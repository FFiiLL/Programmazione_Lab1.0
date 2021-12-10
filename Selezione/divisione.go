package main

import(

   "fmt" 

)

func main() {

    var a,b float64
    
    fmt.Println("inserisci due numeri:")
    fmt.Scan(&a,&b)
    
    
    if b == 0{
        fmt.Println("IMPOSSIBBIRE")
        }else{
        
            fmt.Println("quoziente=",a/b)
        }

}
