package main

import(
  "os"
  "strconv"
  "fmt"
)
func main()  {
  stringa := os.Args[1]
  flag := true

  for _,v := range stringa{

    _, err := strconv.Atoi(string(v))
    if err != nil{
      flag = false
    }
  }


  if flag {
    fmt.Println(Sottostringhe(stringa))
  }

}
func Sottostringhe(s string) []string {

  numeri := make([]int,len(s))
  risultato := []string{}
  var stringa string


  for i,v := range s{
      numeri[i], _ = strconv.Atoi(string(v))
  }


  for i:=0 ; i<len(numeri)-1 ; i++{
    for j:=i+1 ; j<len(numeri) ; j++{
      if seqOrdinata(numeri[i:j+1]){

        for _ ,v := range numeri[i:j+1]{
          stringa += strconv.Itoa(v)
        }
        risultato = append(risultato, stringa)
        stringa = ""
      }
    }
  }


  return risultato
}
func seqOrdinata(numeri []int) bool{
  flag := true

  for i:=0;i<len(numeri)-1;i++{

    if numeri[i] > numeri[i+1]{
      flag = false
      break
    }

  }
  return flag


}
