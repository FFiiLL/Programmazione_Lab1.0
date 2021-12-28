package main

import(

    "fmt"
    "os"
    "strconv"
)

func main() {

    n , _ := strconv.Atoi(os.Args[1])
    fmt.Println(Fattoriali(n))

}



func Fattoriale(n int) int {

    if n == 0 {
        return 1
    }
    return Fattoriale(n-1)*n

}

func Fattoriali(n int) (f []int) {

    for i:=1 ; i<=n ; i++{
    
        f = append(f,Fattoriale(i))
    
    }

    return f
    
}
