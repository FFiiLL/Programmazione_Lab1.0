package main

import(
    
    "fmt"
    "bufio"
    "os"
    

)

func LeggiTesto() string{
    var stringa string
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
    
        if scanner.Text() != ""{
            stringa += scanner.Text() + "\n"
        }else{
            break
        }
    
    }
    
    return stringa[:len(stringa)-1]

}


func InvertiStringa(s string) string{

    var stringa string
    
    for i:=len(s)-1;i>=0;i--{
    
        stringa += string(s[i])
        
    }

    return stringa
}

func main() {

    testo := LeggiTesto()
    testoInvertito := InvertiStringa(testo)
    fmt.Println()
    fmt.Println("Testo invertito:")
    fmt.Println(testoInvertito)
}
