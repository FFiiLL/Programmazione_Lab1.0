package main

import(

    "fmt"
    "os"

)

func main() {

    f, err := os.Create("output.txt")  
    
    if err != nil {
        fmt.Printf("errore durante creazione: %v\n", err)
        return
    }
    
    defer f.Close()
    
   
   
   for i:=0;i<4;i++{
       
       _, err = fmt.Fprintf(f, "Stringaaaaz%d\n", i)
       
       if err != nil {
            fmt.Printf("errore durante scrittura: %v\n", err)
            return
       }
       
   }
}
