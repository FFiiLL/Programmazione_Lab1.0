package main
import (
  "fmt"
  "os"
)

func palindroma(caratteri []rune) bool{

  for i:=0;i<len(caratteri)/2;i++{
    if caratteri[i] != caratteri[len(caratteri)-1-i] {
      return false
    }
  }
  return true
}
func stampaPal(caratteri []rune) {
  var stampa []string
  for i:=0; i<len(caratteri); i++{
    for j:=i+1; j<len(caratteri); j++{
        if palindroma(caratteri[i:j+1]){
          stampa = append(stampa,string(caratteri[i:j+1]))

        }

    }
  }
  fmt.Println(stampa)
}

func main(){
  stringa := os.Args[1]
  caratteri := []rune(stringa)
  stampaPal(caratteri)
}
