package main

import(

   "bufio"
   "os"
   "fmt"
   "strings"
  "unicode"
)

func LeggiTesto()(testo string){
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        if scanner.Text() != ""{
            testo += scanner.Text() + " "
        }
    }
    
    return
}

func SeparaParole(testo string) (parole []string){
    
    testoSeparato := []string{}
    
    testoSeparato = strings.Split(testo, " ")
    testoSeparato = testoSeparato[:len(testoSeparato)-1]
   

    for _, parola := range testoSeparato{
        
        flag := true
        if !unicode.IsLetter(rune(parola[len(parola)-1])){
            parola = parola[:len(parola)-1]
        }
        for _,lettera := range parola{
            if !unicode.IsLetter(rune(lettera)){
                flag = false
            }
        }
        if flag{
            parole = append(parole,parola)
        }
        
    }
    
    return
}

func ContaRipetizioni(parole []string) map[string]int{
    ripetizioni := map[string]int{}
    for _,parola := range parole{
        ripetizioni[parola]++
    }

    return ripetizioni
}

func main() {
    
    testo := LeggiTesto()
    parole := SeparaParole(testo)
    ripetizioni := ContaRipetizioni(parole)
    
    fmt.Println("parole distinte:",len(ripetizioni))
    for k,v := range ripetizioni{
        fmt.Println(k,":",v)
    }
    

}
