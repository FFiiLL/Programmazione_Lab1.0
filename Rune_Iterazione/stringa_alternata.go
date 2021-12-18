package main

import(

    "fmt"
    "math"

)
func IsValidLetter(c rune) bool{

    if (c >='A' && c<='Z') || (c>='a' && c<='z'){
	    return true
    }else{
	    return false
    }  

}

func StringheAlternate(s1 string, s2 string) string{
    lunghezza := math.Max(float64(len(s1)),float64(len(s2)))
    var stringa string
    for i:=0;i<int(lunghezza);i++ {
      
              if i < len(s1){
              stringa += string(s1[i])
             }else{
             stringa += "-"
             }
      
            if i < len(s2){
              stringa += string(s2[i])
             }else{
             stringa += "-"
             }
              
            
       
    }
    
    
   return stringa 
}

func main() {

    var s1,s2 string
    fmt.Scan(&s1)
    fmt.Scan(&s2)
    fmt.Println(StringheAlternate(s1,s2))

}

