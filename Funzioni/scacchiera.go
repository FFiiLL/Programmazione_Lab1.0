package main

import(

    "fmt"

)

func StampaRigaInizioAsterisco(lung int) {
    flag := true
    for i:=0;i<lung;i++{
       if flag{
            fmt.Print("*")
       }else{
            fmt.Print("+")
       }
       flag = !flag
    }
    
}
func StampaRigaInizioPiu(lung int) {
    flag := true
    for i:=0;i<lung;i++{
       if flag{
            fmt.Print("+")
       }else{
            fmt.Print("*")
       }
       flag = !flag
    }
}

func StampaScacchiera(dimensione int) {

    flag := true
    for i:=0;i<dimensione;i++{
       if flag{
            StampaRigaInizioAsterisco(dimensione)
            fmt.Println()
       }else{
            StampaRigaInizioPiu(dimensione)
            fmt.Println()
       }
       flag = !flag
    }



}


func main() {
    var n int
    fmt.Print("Inserisci la dimensione: ")
    fmt.Scan(&n)
    StampaScacchiera(n)

}
