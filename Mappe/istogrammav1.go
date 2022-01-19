package main

import(

    "fmt"
    "os"
    "bufio"
    "unicode"
    "sort"
    "strings"
)

func LeggiTesto() (testo string){
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        testo += scanner.Text() + "\n"
    }
    
    return
}

func Occorrenze(s string) (mappa map[rune]int){
    mappa = make(map[rune]int)
    for _, ch := range s {
        if unicode.IsLetter(ch){
            mappa[ch]++
        }
    }
    
    return
    
}

func StampaIstogramma(occorrenze map[rune]int){
    
    lettereDaOrdinare := []string{}
    
    for carattere := range occorrenze{
        lettereDaOrdinare = append(lettereDaOrdinare,string(carattere))
    }
    sort.Strings(lettereDaOrdinare)
    
    
    for _,chiave := range lettereDaOrdinare{
        
        var runes []rune = []rune(chiave)
        var runa rune = runes[0]
        
        fmt.Print(string(runa),": ",strings.Repeat("*",occorrenze[runa]))
        fmt.Println()
    
    
    }
    
    
    
}

func main() {

    testo := LeggiTesto()
    
    mappa := Occorrenze(testo)
    
    StampaIstogramma(mappa)
    
    
}
