package main

import(

    "fmt"
    "os"
    "strconv"

)

func main() {

    stringa := os.Args[1]
    lunghezza := len(stringa)
    
    
    
    
    for i := 0; i<lunghezza-1;i++{
        prima, _ := strconv.Atoi(string(stringa[i]))
        seconda,_ := strconv.Atoi(string(stringa[i+1]))
        if (prima) < (seconda){
            fmt.Print(string(stringa[i]))
        }
        
    }
    fmt.Println()
    
    
}
