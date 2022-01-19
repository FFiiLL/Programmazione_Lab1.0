package main

import(

    "fmt"
    "os"
    "strings"
    "bufio"
    "sort"
    "strconv"
)

func CreaSequenzaInt (s string) (numeri []int){
    fileLista, _ := os.Open(s)
    
    defer fileLista.Close()
    var stringaLista string
    fmt.Fscanf(fileLista,"%s", &stringaLista)
    
    //sequenze []string
    sequenze := strings.Split(stringaLista,";")
    
    
    for _, sequenza := range sequenze{
        fmt.Println(sequenza)
        
        fileSequenza, _ := os.Open(sequenza)
        defer fileSequenza.Close()
        var stringaSequenza string
        
        scanner := bufio.NewScanner(fileSequenza)
        
        for scanner.Scan(){
        
            stringaSequenza += scanner.Text()
        
        }
        fmt.Println(stringaSequenza)

        sliceStringheInteri := strings.Split(stringaSequenza," ")
        
        for _,strInt := range sliceStringheInteri{
            intero, err := strconv.Atoi(strInt)
            if err != nil{
                break
            }
            fmt.Println(intero)
            numeri = append(numeri,intero)
        }
        
    }
   
    sort.Ints(numeri)
    return
}

func main() {
    
    
    numeri := CreaSequenzaInt(os.Args[1])
   
    seqRes,_ := os.Create("sequenza_risultante.txt")
    defer seqRes.Close()
    for _, numero := range numeri{
        fmt.Fprintf(seqRes,"%d ",numero)
    }
    fmt.Fprintln(seqRes)
     /*   
    //-------------------------------------------
    fileLista, _ := os.Open(os.Args[1])
    
    defer fileLista.Close()
    var stringaLista string
    fmt.Fscanf(fileLista,"%s", &stringaLista)
    
    //sequenze []string
    sequenze := strings.Split(stringaLista,";")
    
    var numeri []int
    for _, sequenza := range sequenze{
        
        fileSequenza, _ := os.Open(sequenza)
        defer fileSequenza.Close()
        var stringaSequenza string
        
        scanner := bufio.NewScanner(fileSequenza)
        
        for scanner.Scan(){
        
            stringaSequenza += scanner.Text()
        
        }
        

        sliceStringheInteri := strings.Split(stringaSequenza," ")
        
        for _,strInt := range sliceStringheInteri{
            intero, _ := strconv.Atoi(strInt)
            numeri = append(numeri,intero)
        }
        
    }
   
    sort.Ints(numeri)
    //-------------------------
    */
    
    
    
}
