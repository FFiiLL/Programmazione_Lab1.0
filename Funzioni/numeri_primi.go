package main

import(

    "fmt"

)

func EPrimo(n int) bool {

    flag := true
    
    for i:=2;i<n/2;i++{
        
        if n % i == 0{
            flag = false
            return flag
        }
    
    
    }

    return flag

}
func NumeriPrimi(limite int) {

    for i := 2 ; i<limite ; i++{
    
        if EPrimo(i){
            fmt.Print(i," ")
        }
    }



}


func main() {

   var n int
   fmt.Print("Inserisci un numero: ")
   fmt.Scan(&n)
   if n > 0{
       fmt.Print("numeri primi inferiori a ",n,"\n")
       NumeriPrimi(n)
       fmt.Println()
   }else{
    fmt.Println("numero nullo o negativo")
    }
}
