package main

import(

  "fmt"
  "os"  
  "math"
  "strconv"
)

func main() {

    f, err := os.Create("potenze.txt")  
    
    if err != nil {
        fmt.Printf("errore durante creazione: %v\n", err)
        return
    }
    
    defer f.Close()
    
    base,_:= strconv.Atoi(os.Args[1])
    esp,_ := strconv.Atoi(os.Args[2])
    
    for i:=0;i<=esp;i++{
       
       _, err = fmt.Fprintf(f, "%.0f ", math.Pow(float64(base),float64(i)) )
      
       if err != nil {
            fmt.Printf("errore durante scrittura: %v\n", err)
            return
       }
       
   }
    _, err = fmt.Fprintf(f,"\n")
}

