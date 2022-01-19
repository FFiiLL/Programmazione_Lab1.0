package main

import(

   "fmt" 

)

func main() {
    ///[righe][colonne]
    var matrix = [3][2]int{{0,1},{2,3},{4,5}}
   //stampa solo le righe
    for _, riga := range matrix{
        fmt.Println(riga)
    
    
    }
    //stampa gli elementi distinti 
    for _, riga := range matrix{
        for _, elem := range riga{
            fmt.Print(elem," ")
        }
        fmt.Println()
    
    }   
   
   
   //stampa alternativa
    for i:=0 ; i<len(matrix) ; i ++{
    
        for j:=0 ; j<len(matrix[i]); j++{
            fmt.Print(matrix[i][j]," ")
        }
        fmt.Println()
    }
    
    
    
    
    
    
    fmt.Println("-----------")
    fmt.Println("-----------")
    
    
    //SLICE permettono righe e colonne di lunghezza variabiale
   
    
    
     ///[righe][colonne]
    var slice2d [][]int = [][]int{{0,1},{2},{4,5,6,7}}
    for _, riga := range slice2d{
        fmt.Println(riga)
    
    
    }
    fmt.Println("-----------")
    slice2d[0] = append(slice2d[0],0)
    slice2d = append(slice2d, []int{8,9,10})
   //stampa solo le righe
    for _, riga := range slice2d{
        fmt.Println(riga)
    
    
    }
    fmt.Println("-----------")
    //stampa gli elementi distinti 
    for _, riga := range slice2d{
        for _, elem := range riga{
            fmt.Print(elem," ")
        }
        fmt.Println()
    
    }   
   
   fmt.Println("-----------")
   //stampa alternativa
    for i:=0 ; i<len(slice2d) ; i ++{
    
        for j:=0 ; j<len(slice2d[i]); j++{
            fmt.Print(slice2d[i][j]," ")
        }
        fmt.Println()
    }
    
    
    
    
    
    
    
    
    
}
