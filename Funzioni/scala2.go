package main

import(

    "fmt"   

)

func StampaGradino(gradino int){

    if gradino == 0{
    
        fmt.Println("***")
        fmt.Println("*  ")
    
    }
    if gradino > 0{
    
        for i:=0;i<gradino*2;i++{
            fmt.Print(" ")
        }
        fmt.Println("***")
        for i:=0;i<gradino*2;i++{
            fmt.Print(" ")
        }
        fmt.Println("*  ")
    
    }
     
}

func StampaScala(gradini int){

    for i:=gradini;i>0;i--{
    
        StampaGradino(i)
    
    }

}


func main() {

    var n int
    fmt.Scan(&n)
    StampaScala(n)
}
