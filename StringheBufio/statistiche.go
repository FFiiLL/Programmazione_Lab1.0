package main

import(

    "os"
    "bufio"
    "fmt"
    "strings"
    "unicode"
)

func LeggiTesto() string {
    var stringa string
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        
        stringa += scanner.Text() + "\n"

    }
    
    return stringa[:len(stringa)-1]
}

func StatisticheParole(s string)(int,int,int){
     linee := strings.Split(s, "\n")
        var numParole int
        var lunghezzaTotale int
        for _, linea := range linee{
            
            parole := strings.Split(linea, " ")
            numParole += len(parole)
            for _, parola := range parole{
                
                lung := 0
                
                for _,ch := range parola{
                    if unicode.IsLetter(ch){
                        lung++
                    }
                }
                
                lunghezzaTotale += lung
            }
      }
      
      return len(linee),numParole, lunghezzaTotale

}


func main() {

    testo := LeggiTesto()
    fmt.Println()
     

   
    nLinee, nParole, lParole := StatisticheParole(testo)
    
    fmt.Println("numero linee:",nLinee)
    fmt.Println("numero parole:",nParole)
    fmt.Println("lunghezza media:",float64(lParole)/float64(nParole))
    

}
