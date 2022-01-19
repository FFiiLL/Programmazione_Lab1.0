package main

import(

    "fmt"
    "os"
    "bufio"
    "strconv"
    "unicode"
)

func LeggiTesto() string {
    
    var testo string
    scanner := bufio.NewScanner(os.Stdin)
    
    if scanner.Scan(){
    testo += scanner.Text()
    }
    return testo
}

func NumeroNascosto(testo string) (int, error) {
    var stringaNumeri string
    for _, ch := range testo{
    
         if unicode.IsDigit(ch){
            stringaNumeri += string(ch)
         }
         
         
        
    }
    numero, err := strconv.Atoi(stringaNumeri)
    return numero, err

}

func main() {

    stringa := LeggiTesto()
    numero,err := NumeroNascosto(stringa)
    if err == nil{
        fmt.Println(numero*2)
    }else{
        fmt.Println("nessun numero nascosto")
    }
    
}
