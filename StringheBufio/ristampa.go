package main

import(

    "fmt"
    "bufio"
    "os"

)

func LeggiTesto() string {

    var stringa string
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        
        stringa +=  scanner.Text() + "\n"
        
    }
    
    return stringa
}

func main() {
    fmt.Println("inserisci testo(termina con CTRL+D)")
    testo := LeggiTesto()
    fmt.Println("\ntesto letto:")
    fmt.Println(testo)

}
