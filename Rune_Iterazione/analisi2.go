package main

import(

 "fmt"   

)

func IsValidLetter(c rune) bool{

    if (c >='A' && c<='Z') || (c>='a' && c<='z'){
	    return true
    }else{
	    return false
    }  

}

func IsUpperCase(c rune) bool{
    
    if (c >='A' && c<='Z') {
	    return true
    }else{
	    return false
    }
}

func IsVocale(c rune) bool{
    
    if c == 'A' || c == 'E' ||c == 'I' ||c == 'O' ||c == 'U' ||c == 'a' ||c == 'e' ||c == 'i' ||c == 'o' ||c == 'u' {
        return true
    }else{
        return false
    }
    
}
func main() {

   var s string
   fmt.Scan(&s) 
   var maiuscoleVocali, minuscoleVocali, maiuscoleConsonanti, minuscoleConsonanti int
   
   for _, carattere := range s{
    
        if IsValidLetter(carattere){
            if IsUpperCase(carattere){
               if IsVocale(carattere){
                maiuscoleVocali++
               }else{
                maiuscoleConsonanti++
               }
            }else{
               if IsVocale(carattere){
                minuscoleVocali++
               }else{
                minuscoleConsonanti++
               } 
            }
            
           
        } 
   }
   
   fmt.Printf("Maiuscole vocali: %d\nMinuscole vocali: %d\nmaiuscole consonanti: %d\nminuscole Consonanti: %d\n",maiuscoleVocali, minuscoleVocali, maiuscoleConsonanti, minuscoleConsonanti)
}
