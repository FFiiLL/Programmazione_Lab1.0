package main

import(

    "fmt"
    "os"
    "sort"
    "strconv"
)

func Primo(n int) bool{

    flag := true
    
    for i:=2;i<n;i++{
        if n%i == 0{
            flag = false
        }
    }
    return flag
}


func Primi(stringa string) (primi []int){
    lungh := len(stringa)
    for i:=0 ; i<3 ; i++{
        
        numero,_ := strconv.Atoi(stringa[i:lungh])
        fmt.Println(numero)
        if Primo(numero){
        
            primi = append(primi,numero)
        
        }
    
    }

    sort.Ints(primi)
    return primi

}

func main() {

    stringaNumero := os.Args[1]
    primi := Primi(stringaNumero)
    fmt.Println(primi)

}
