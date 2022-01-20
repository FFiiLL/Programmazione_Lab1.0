package main

import(

    "os"
    "fmt"
    "strconv"
    "sort"
)

func ePrimo(n int) bool{
    flag := true
    for i:=2 ; i<n ; i++{
        if n % i == 0{
            flag = false
            return flag
        }
    }
    return flag

}

func main() {

    stringa := os.Args[1]
    var numeri []int
    for i:=0;i<len(stringa)-1;i++{
        for j:=i;j<len(stringa);j++{
        
            numero,_ := strconv.Atoi(stringa[i:j+1])
            numeri = append(numeri,numero)
        }
    }
    
    sort.Ints(numeri)
    
    
    var primi map[int]bool
    primi = make(map[int]bool)
    
    for _, numero := range numeri{
        
        primi[numero] = ePrimo(numero)
        
    }
    
    
    var numeriPrimi []int
    
    for k,v := range primi{
    
        if v && k<=100 && k>1{
            numeriPrimi = append(numeriPrimi,k)
        }
    
    }
    
    sort.Ints(numeriPrimi)
    
    for i:=len(numeriPrimi)-1 ; i>=0 ; i--{
        if i != 0{
            fmt.Printf("%d, ",numeriPrimi[i])
        }else{
            fmt.Printf("%d\n",numeriPrimi[i])
        }
    }
    
}
