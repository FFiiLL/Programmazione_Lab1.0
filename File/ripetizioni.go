package main

import(

    "os"
    "bufio"
    "fmt"
   "strings"
  "unicode"

)

func main() {

    fIn, _ := os.Open(os.Args[1])
    fOut, _ := os.Create("statistiche.txt")
    defer fIn.Close()
    defer fOut.Close()

    var testo string
    scanner := bufio.NewScanner(fIn)
    
    for scanner.Scan(){
        testo += scanner.Text() 
    } 
    
 
    
    parole := strings.Split(testo," ")
    for i,v := range parole{
        fmt.Printf("%d %s\n",i,v)
    }
    
    
    var ripetizioni map[string]int
    ripetizioni = make(map[string]int)
    for _, parola := range parole{
            for _, lettera := range parola{
                
                if unicode.IsSpace(lettera){
                    continue
                }
            }
            
            
            ripetizioni[strings.ToUpper(parola)]++
    }
    fmt.Fprintf(fOut,"parole uniche %d\n",len(ripetizioni))
    for k,v := range ripetizioni{
        fmt.Fprintf(fOut,"%s - %d\n",strings.ToUpper(k),v)
    }
    
    
}
