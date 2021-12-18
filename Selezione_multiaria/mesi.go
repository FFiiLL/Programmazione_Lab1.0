package main

import(

    "fmt"

)

func numeroDiGiorni(s string) int {
    var n int
    switch s{
        
        case "novembre":
            n = 30
        case "aprile":
            n = 30
        case "giugno":
            n = 30
        case "settembre":
            n = 30            
        case "febbraio":
            n = 28
        case "gennaio":
            n = 31
        case "marzo":
            n = 31
        case "maggio":
            n = 31
        case "luglio":
            n = 31
        case "agosto":
            n = 31
        case "ottobre":
            n = 31
        case "dicembre":
            
            n = 31
        default:
            fmt.Println("nome non corretto")
            n = 0
    }    
    
        return n
}

func main() {

    var mese string
    
    fmt.Scan(&mese)
    x := numeroDiGiorni(mese) 
    if x!=0 {
     fmt.Println(x)
    }
    
    
    

}

