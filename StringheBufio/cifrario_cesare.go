package main

import(

    "bufio"
    "os"
    "fmt"
    

)

func LeggiNumero() (k int){
    fmt.Scanln(&k)
    return
}

func LeggiTesto() (testo string){
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        testo += scanner.Text() + "\n"
    }
    
    return
}

func CifraCarattere(c rune, chiave int) (r rune){
    
    if c>='a' && c<='z' || c>='A' && c<='Z'{
        r = c+rune(chiave)
    }else{
        r = c
    }
    
    return
}

func CifraTesto(s string, chiave int) (cifrato string){
    
    for _, ch := range s{
        cifrato += string(CifraCarattere(ch,chiave)) 
    }
    
    return
}
func main() {
    fmt.Print("inserisci numero:")
    k := LeggiNumero()
    fmt.Println("inserisci testo su piu righe termina con CTRL+D------------")
    testo := LeggiTesto()
    fmt.Println("testo cifrato:")
    fmt.Println(CifraTesto(testo,k))

}
