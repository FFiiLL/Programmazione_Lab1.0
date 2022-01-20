package main

import (
  "fmt"
  "os"
)


func isBalanced(stringa string) bool{


  sequenza := make([]rune,0,len(stringa))

  for _, r := range stringa {

    if r == '('{
      sequenza = append(sequenza, r)
    }else{
      if len(sequenza) == 0{
        return false
      }
      sequenza = sequenza[:len(sequenza)-1]
    }

  }

  if len(sequenza) != 0{
    return false
  }
  return true


}

func stampaSottosequenze(stringa string) (bool, []string) {

  flag := false
  seq := []rune(stringa)
  risultato := make([]string,0, len(stringa))


  for i:=0 ; i<len(stringa)-1;i++{
    for j:=i+1 ; j<len(stringa);j++{
      if isBalanced(string(seq[i:j+1])){
      //  fmt.Println(string(seq[i:j+1]))
        risultato = append(risultato,string(seq[i:j+1]) )

        flag = true
      }
    }
  }
  return flag, risultato
}




func main() {

  stringa := os.Args[1]


  for _, carattere := range stringa{
    if carattere != '(' && carattere != ')'{
      fmt.Println("input con caratteri diversi dalle parentesi tonde")
      return
    }
  }


if isBalanced(stringa){
  fmt.Println("bilanciata")

}else{
  fmt.Println("non bilanciata")
}

flag, risultato := stampaSottosequenze(stringa)

if !flag{
    fmt.Println("nessuna sotto sequenze bilanciat:")
  }else{
    fmt.Println("sotto sequenze bilanciat:")
    for _ , v := range risultato{
      fmt.Println(v)
    }
  }
}
