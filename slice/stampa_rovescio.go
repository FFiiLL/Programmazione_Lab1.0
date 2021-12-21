package main

import(

    "fmt"

)

func main() {

    var n int
    fmt.Scan(&n)
    var sl []int   
    sl = make([]int,n)
    
    fmt.Println(sl)
    for i:=0 ; i<len(sl) ; i++ {
        
      fmt.Scan(&sl[i])
    }
    
    fmt.Println(sl)
    
    for i := range sl{
    
        fmt.Print(sl[n-1-i]," ")
    
    } 

}
