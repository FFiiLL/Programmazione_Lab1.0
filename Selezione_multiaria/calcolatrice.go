package main

import(

   "fmt" 

)

func main() {

    var a,b float64
    var op rune
    
    fmt.Scanf("%f %c %f",&a,&op,&b)
    switch op{
        
        case '+':
            fmt.Print("Risultato: ",a+b)
            fmt.Println()
        case '-':
            fmt.Print("Risultato: ",a-b)
            fmt.Println()
        case '*':
            fmt.Print("Risultato: ",a*b)
            fmt.Println()
        case '/':
            fmt.Print("Risultato: ",a/b)
            fmt.Println()
        default:
            fmt.Println("operatore non riconosciuto")
    }

}

