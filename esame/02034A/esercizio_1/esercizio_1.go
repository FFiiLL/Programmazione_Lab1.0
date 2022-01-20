package main

import(

    "fmt"
    "strings"
    "strconv"
)

func isPerpend(m1,m2 float64) bool{

    if m1 - (-1/m2) < 0.0001{
        return true
    }else{
        return false
    }
    


}

func main() {

    var retta1,retta2  string
    fmt.Scanln(&retta1, &retta2)
    
    
    slR1 := strings.Split(retta1,"x")
    a1 := slR1[0]
    retta1 = slR1[1]
    slR1 = strings.Split(retta1,"y")
    b1 := slR1[0]
    
    
    slR2 := strings.Split(retta2,"x")
    a2 := slR2[0]
    retta2 = slR2[1]
    slR2 = strings.Split(retta2,"y")
    b2 := slR2[0]
    A1,_ := strconv.Atoi(a1)
    B1,_ := strconv.Atoi(b1)
   
    A2,_ := strconv.Atoi(a2)
    B2,_ := strconv.Atoi(b2)
    var m1,m2 float64
    m1 = (-1*float64(A1))/float64(B1)
    m2 = (-1*float64(A2))/float64(B2)
    
    if isPerpend(m1,m2){
    
        fmt.Println("le due rette sono perpendicolari")
    
    }else{
         fmt.Println("le due rette non sono perpendicolari")
    }
    
   
   
}
