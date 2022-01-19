package main

import(

    "fmt"
    "os"
    "strings"
    "strconv"
    "math"

)

func main() {
//SEPARA A B C 
    funzione := os.Args[1]
    soglia ,_ := strconv.ParseFloat(os.Args[2],64)
    epsilon ,_ := strconv.ParseFloat(os.Args[3],64)
    
    
    
    
    slice := strings.Split(funzione,"+")    
    funzione = ""
    for i:=0;i<len(slice);i++{
       funzione += string(slice[i]) 
    }    
    slice = strings.Split(funzione,"x^2")
     a,_ := strconv.Atoi(slice[0])
    funzione = ""
    for i:=1;i<len(slice);i++{
       funzione += string(slice[i]) 
    }   
    slice = strings.Split(funzione,"x")
     b,_ := strconv.Atoi(slice[0])
    funzione = ""
    for i:=1;i<len(slice);i++{
       funzione += string(slice[i]) 
    }
    slice = strings.Split(funzione,"=0")
    c,_ := strconv.Atoi(slice[0])
    
   // CALCOLA  
    
     x1 := ((float64(b)*-1) + math.Sqrt(math.Pow(float64(b),2) - (4 * float64(a) * float64(c)))) / (2 * float64(a))
     x2 := ((float64(b)*-1) - math.Sqrt(math.Pow(float64(b),2) - 4 * float64(a) * float64(c))) / (2 * float64(a))
     
    // ARROTONDO A 3 DECIMALI
    x1 = (math.Round(x1*1000))/1000
    x2 = (math.Round(x2*1000))/1000
     
    if  x1 != x2 && !math.IsNaN(x1){
        fmt.Printf("Esistono due soluzioni reali: %.3f e %.3f\n", x1,x2)
    }else if x1 == x2{
         fmt.Printf("Esiste una soluzione reale: %.3f \n ", x1)
    }else{
        fmt.Print("NON Esistono soluzioni reali\n")
    }
}
