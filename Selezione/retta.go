package main

import(

   "fmt" 

)

func main() {
    var m, q, x, y float64
    fmt.Print("inserisci m e q: ")
    fmt.Scan(&m,&q)
    fmt.Print("inserisci x e y: ")
    fmt.Scan(&x,&y)
    
    if y == m * x + q{
        fmt.Println("il punto appartiene alla retta")
     }else if y > m * x + q {
     
        fmt.Println("il punto sta sopra")
     
     }else {fmt.Println("il punto sta sotto")}   
}
