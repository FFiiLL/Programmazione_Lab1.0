package main

import(

    "os"
    "fmt"
    "strconv"
)

func main() {
    prodotto := 1
    
    for i:=1 ; i<len(os.Args) ; i++{
        
        if n, err := strconv.Atoi(os.Args[i]) ; err == nil{
            prodotto*=n
        }
        
    }
    
    fmt.Println(prodotto)

}
