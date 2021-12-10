package main

import(

    "fmt"

)

func main() {
    var voto int 
    
    fmt.Print("inserisci il voto: ")
    fmt.Scan(&voto)
    
    if voto<0 || voto>100{
        fmt.Println("ERRORE")
    }else if voto<60{
        fmt.Println("insufficiente")
        }else if voto < 70{
            fmt.Println("sufficiente")
            }else if voto < 80{
                fmt.Println("buono")
                }else if voto < 90{
                    fmt.Println("distinto")
                    }else if voto <= 100{
                        fmt.Println("ottimo")
                        }
    

}
