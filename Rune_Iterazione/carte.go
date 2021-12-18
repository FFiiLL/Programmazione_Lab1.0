package main

import(

    "fmt"   

)

func main() {

    var carta rune = 127153
    for i:=0;i<10;i++{
    
       fmt.Println(string(carta),carta)
       carta++
    }

}
