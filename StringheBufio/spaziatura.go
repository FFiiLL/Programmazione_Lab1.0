package main

import(

    "bufio"
    "os"
    "fmt"
   
    "unicode"
    

)

func LeggiTesto() (testo string){

    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        testo += scanner.Text() + "\n"
    }
    
    return
}

func Spazia(s string) (spazi string) {
    for i, ch := range s{
        if !unicode.IsSpace(ch){
            if unicode.IsSpace(rune(s[i+1])){
                spazi += string(ch) 
            }else{
                spazi += string(ch) + " "
            }
           
        }else{
            spazi += string(ch)
        }
        
    }
    
    return
}

func main() {
    fmt.Println("inserisci testo su piu righe termina con CTRL+D------------")
    testo := LeggiTesto()
    fmt.Println("risultato-----------------------------------------------")
    spazi := Spazia(testo)
    fmt.Println(spazi)
    

}
