package main

import (
    "fmt"
    "os"
    "bufio"
    //"io"
)

func main(){

    f, err := os.Open("file.txt")
    
    if err != nil {
        fmt.Printf("Errore durante l'apertura: %v\n", err)
        return
    }

    defer f.Close()
    
    scanner := bufio.NewScanner(f)
    
    for scanner.Scan() {
    
                   // var stringa string
                  //  var x,y float64
                
                    
                  //  _, err := fmt.Fscan(f, &stringa, &x, &y)
                    
                   // if err == io.EOF{
                    //    break
                   // }
       var linea string
       linea = scanner.Text()
       err := scanner.Err()
       if err != nil {
           fmt.Printf("Errore lettura file: %v\n", err) 
           return
       }
        
       fmt.Println(linea)
    }
    
}
