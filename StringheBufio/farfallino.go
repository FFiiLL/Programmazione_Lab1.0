package main

import(

    "fmt"
    "strings"
    "bufio"
    "os"

)

func LeggiTesto() (testo string){
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        testo += scanner.Text() + "\n"
    }
    return
}

func TraduciCarattereInFarfallino(c rune) (farfallino string) {
   farfallino += string(c)
   
   if strings.ContainsRune("aeiouàèéìòù", c){
        farfallino += "f" + string(c)
   } 
   if strings.ContainsRune("AEIOUÀÁÈÉÌÒÙ", c){
        farfallino += "F" + string(c)
   } 
   return
}

func TraduciTestoInFarfallino(s string) (testo string) {
    
   for _, ch := range s{
        testo += TraduciCarattereInFarfallino(ch)
        
   } 
   return
}
func main() {
    fmt.Println("inserisci un testo su piu righe fine CTRL D")
    testo := LeggiTesto()
    farf := TraduciTestoInFarfallino(testo)
    fmt.Println("Risultato:")
    fmt.Print(farf)

}
