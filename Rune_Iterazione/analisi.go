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
   var maiuscole, minuscole, consonanti, vocali int
   
   for _, carattere := range s{
    
        if IsValidLetter(carattere){
            if IsUpperCase(carattere){
               maiuscole++
            }else{
                minuscole++
            }
            
            if IsVocale(carattere){
                vocali++
            }else{
                consonanti++
            }
        } 
   }
   
   fmt.Printf("Maiuscole: %d\nMinuscole: %d\nVocali: %d\nConsonanti: %d\n",maiuscole,minuscole,vocali,consonanti)
}
