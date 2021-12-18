package main

import(

    "fmt"
    "math"
)

func IsTernaPitagorica( a int, b int, c int) bool {
 
    if math.Pow(float64(c),2) == math.Pow(float64(a),2) + math.Pow(float64(b),2){
        return true
    }else{
        return false
    }
    
}

func TernePitagoriche(limite int) {
    
    for i:=1;i<limite-2;i++{
        for j:=1;j<limite-1;j++{
            for k:=1; k<limite;k++{
                if IsTernaPitagorica(i,j,k){
                    fmt.Printf("(%d,%d,%d)\n",i,j,k)
                }
            }
            
        }
    
    }   
    
    
}
func main() {

    var n int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    if n <= 0 {
        fmt.Println("soglia non positivo")
    }else{
        fmt.Println("TERNE: ")
        TernePitagoriche(n)
        
    }

}
