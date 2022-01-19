package main

import(

    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {

    file, err := os.Open(os.Args[1])
    if err != nil{
        fmt.Println("error %v",err)
        return
    }
    defer file.Close()
    pattern := os.Args[2]
    
    scanner := bufio.NewScanner(file)
    var testo []string
    for scanner.Scan() {
    
        testo = append(testo,scanner.Text()) 
    
    
    }
    
   
    for _, riga := range testo{
        if strings.Contains(riga,pattern){
            fmt.Println(riga)
        }
    }
}
