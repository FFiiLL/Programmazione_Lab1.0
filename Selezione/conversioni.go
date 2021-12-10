package main

import(

    "fmt"

)

func main() {
    var selezione int
    var valore float64
    fmt.Print("Scegli la conversione:\n1) secondi -> ore\n2) secondi -> minuti\n3) minuti -> ore\n4) minuti -> secondi\n5) ore -> secondi\n6) ore -> minuti\n7) minuti -> giorni e ore\n8) minuti -> anni e giorni\n")
    fmt.Scan(&selezione)
    
    if selezione<1||selezione>8{
        fmt.Println("scelta errata")
    }else{
    
        if selezione == 1{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"secondi corrispondo a",valore/3600,"ore")
            
        }else if selezione == 2{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"secondi corrispondo a",valore/60,"minuti")
        
        }else if selezione == 3{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"minuti corrispondo a",valore/60,"ore")
        
        }else if selezione == 4{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"minuti corrispondo a",valore*60,"secondi")
        
        }else if selezione == 5{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"ore corrispondo a",valore*3600,"secondi")
        
        }else if selezione == 6{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            fmt.Println(valore,"ore corrispondo a",valore*60,"minuti")
        
        }else if selezione == 7{
            fmt.Print("inserisci valore da convertire: ")
            fmt.Scan(&valore)
            var giorni int
            
            giorni = int(valore)/(60*24)
            
            fmt.Println(valore,"minuti corrispondo a",giorni,"giorni",valore/60-(float64(giorni) *  24),"ore")
        
        }else if selezione == 8{
            fmt.Print("inserisci valore da convertire: ")
                fmt.Scan(&valore)
                var anni int
                
                anni = (((int(valore)/60)/24)/365)
                
                var giorni int = ((int(valore)/60)/24) - anni * 365
                
                fmt.Println(valore,"minuti corrispondo a",anni,"anni",giorni,"giorni")
        
        }
        
        
        
     }
}
