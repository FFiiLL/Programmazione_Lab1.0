package main

import(

    "bufio"
    "os"
    "fmt"
    "strings"

)

func LeggiTesto() (testo string){
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        if scanner.Text() == "" {
            break
        }else{
            testo += scanner.Text() + "\n"
        }
    }
    return
}

func TrasformaCarattere(c rune) rune{
    
    var runa rune
    
    if strings.ContainsRune("aeiou",c){
        runa = 'u'
    }else{
        runa = c
    }
    
    return runa
}

func Garibaldi(s string) string{
    
    var stringa string
    
    for _, ch := range s{
        
        stringa += string(TrasformaCarattere(ch))
        
    }
    return stringa
}

func main() {
    fmt.Println("testi su piu righe------------")
    testo := LeggiTesto()
    fmt.Println("trasformazione:-----------")
    fmt.Println(Garibaldi(testo))

}
