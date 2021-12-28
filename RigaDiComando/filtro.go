package main

import(

    "os"
    "fmt"
    "strconv"

)

func LeggiNumeri() (numeri []int){

    for i:=1 ; i<len(os.Args) ; i++{         
        if n, err := strconv.Atoi(os.Args[i]) ; err == nil && n>=0 && n<=100{
             numeri= append(numeri,n)
        }           
    }
    return
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int){

    for _,voto:=range voti{
        if voto>=60{
            sufficienti = append(sufficienti,voto)
        }else{
            insufficienti = append(insufficienti,voto)
         }
    }
    
    return
}



func main() {

    
    numeri := LeggiNumeri()
    suff,insuff := FiltraVoti(numeri)
    fmt.Println(suff)
    fmt.Println(insuff)
}
