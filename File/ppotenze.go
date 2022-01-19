package main

import(

    "os"
    
    "fmt"
    "strconv"
    "math"
)

func main() {
   
    b, _ :=  strconv.Atoi(os.Args[1])
    e, _ :=  strconv.Atoi(os.Args[2])
    
    file, err := os.Create("ppotenze.txt")
    
    if err != nil{
        fmt.Println("error %v",err)
        return
    }
    defer file.Close()
    for exp:=0 ; exp<=e ; exp++{
        fmt.Fprintf(file,"%1.0f ",math.Pow(float64(b),float64(exp)))
    }
    fmt.Fprintln(file)
}
