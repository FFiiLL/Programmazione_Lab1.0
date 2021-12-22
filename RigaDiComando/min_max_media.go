package main

import(

    "os"
    "fmt"
    "strconv"
    "math"
)

func main() {
   
   
    numeri := LeggiNumeri()
    fmt.Println(numeri)
    fmt.Println("massimo",Massimo(numeri),"\n","minimo",Minimo(numeri))
    fmt.Println("media",Media(numeri))

}




func LeggiNumeri() (numeri []int){


    for i:=1 ; i<len(os.Args) ; i++{
                
                if n, err := strconv.Atoi(os.Args[i]) ; err == nil{
                    numeri= append(numeri,n)
                }
                
    }
 
    return

}

func Massimo(sl []int) int {
       var max int = sl[0]
      for i:=1 ; i<len(sl); i++{
      
        max = int(math.Max(float64(sl[i]),float64(max)))
        
        
      }
        
     return max   
}

func Minimo(sl []int) int {
      var min int = sl[0]
      for i:=1 ; i<len(sl); i++{
      
        min = int(math.Min(float64(sl[i]),float64(min)))
        
      }
        
     return min
}

func Media(sl []int) float64 {

    var somma int
    
    for _, numero := range sl{
    
        somma += numero
        
        
    }
    
    return float64(somma)/float64(len(sl))

}
