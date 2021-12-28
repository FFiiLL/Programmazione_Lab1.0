package main

import(

    "os"
    "fmt"
    "math/rand"
    "time"
    "strconv"
    

)
func Genera(soglia int) []int {

    rand.Seed(int64(time.Now().Nanosecond()))
    
    var numeri []int
    
    numeroGenerato := rand.Intn(101)
    numeri = append(numeri,numeroGenerato)
    
    for numeroGenerato > soglia{
        
       numeroGenerato = rand.Intn(101) 
       numeri = append(numeri,numeroGenerato)
    }
    
    return numeri
}

func main() {
    
    soglia, _ := strconv.Atoi(os.Args[1])
    numeri := Genera(soglia)
    fmt.Println(numeri)
    
    for _,numero := range numeri{
        if numero > soglia{
            fmt.Print(numero," ")
        }
    }
    fmt.Println()
} 
