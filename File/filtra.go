package main

import(
    "os"
    "fmt"
    "bufio"
    "strings"
)

func Errore(err error) bool {

    if err != nil{
        
        fmt.Printf("Errore: %v\n", err)
        return true
    }
    
    return false
    
}

func main() {

    f, err := os.Open(os.Args[1])
    
    if Errore(err) {
       return
    }
    defer f.Close()
    pattern := os.Args[2]
    
    scanner := bufio.NewScanner(f)
    
    for scanner.Scan(){
      
        var stringa string
        
        stringa = scanner.Text()
        
      
        if strings.Contains(stringa,pattern){
          fmt.Println(stringa)
        }
        
    
        
          
        
    }
    
    

}
