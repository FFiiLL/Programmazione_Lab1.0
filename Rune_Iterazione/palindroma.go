package main

import(

    "fmt"

)

func IsPalindroma(s string) bool {
    
    flag := true
    
    for i:=0;i<len(s);i++{
        if s[i] != s[len(s)-1-i]{
            flag = false
        }
    }
    
    
    return flag
}

func main() {

   var s string
   fmt.Scan(&s)
   if IsPalindroma(s){
        fmt.Println("palindroma")
   }else{
        fmt.Println("NON palindroma")
   } 

}
