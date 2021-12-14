package main

import(

    "fmt"

)

func main() {

    var n int
    fmt.Scan(&n)
    //TRIANGOLO DRITTO
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            if i == 0 || j == i ||  j == n-1{
                fmt.Print("*")
            }else{
                fmt.Print(" ")
            }
           
        }
        
         fmt.Println()
         
    }
    //TRIANGOLO ROVESCIO
    for i:=0;i<n;i++{
        for j:=0;j<2*n-1;j++{
            if j == n-1|| (i == n-1 && j >= n) || (j>=n && j == (n -1) + i ) {
                fmt.Print("*")
            }else{
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

}
