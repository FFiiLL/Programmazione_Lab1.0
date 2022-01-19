package main

import(

    "fmt"
    "os"
    "strconv"
    

)

func CreaTavolaPitagorica(n int) [][]int {
    tavola := make([][]int,n)
    
    for i:=0 ; i<n ; i++{
        tavola[i] = make([]int,n)
        for j:=0 ; j<n ; j++{
            tavola[i][j] = (i+1) * (j+1)
        }
    
    }
    
    return tavola
}

func StampaTavolaPitagorica(tavola [][]int){
    for _ ,riga := range tavola{
        for _,numero := range riga{
            fmt.Printf("%3d ",numero)
        }
        fmt.Println()
    }
}

func main() {
    n, _ := strconv.Atoi(os.Args[1])
    tavola := CreaTavolaPitagorica(n)
    StampaTavolaPitagorica(tavola)
    

}
