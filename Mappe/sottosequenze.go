package main

import(

    "fmt"
    "os"
    

)

func main() {

    sequenza := os.Args[1:len(os.Args)]
    
    sottosequenze := []string{}
    sequenzaRune := []rune{}
   
    for _, lettera := range sequenza{
        var runes []rune = []rune(lettera)
        var runa rune = runes[0]
        sequenzaRune = append(sequenzaRune, runa)
    
    }
   
  
   
    for i:=0 ; i<len(sequenzaRune)-2 ; i++{
        for j:= i+2 ; j<len(sequenzaRune) ; j++{
            
            var sottoseq string
            sottoseq = string(sequenzaRune[i:j+1])
            
            if sottoseq[0] == sottoseq[len(sottoseq)-1] && len(sottoseq)>2{
                sottosequenze = append(sottosequenze,sottoseq)
            }
        }
    }
    
    
    
    occorrenze := map[string]int{}
    
    for _,v := range sottosequenze{
        occorrenze[v]++
    }
    for k,v := range occorrenze{
        fmt.Println(k,":",v)
    }
}
