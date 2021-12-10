package main

import(
    "fmt"
)


func main() {

    var cont float64
    var somma float64
    for true {
    
        var n float64
        fmt.Scan(&n)
        if n > 0{
            somma += n
            cont ++
        }else{
            break
          }
    
    
    }
    
    
    
    fmt.Println("Media aritmetica:", (somma/cont))

}
