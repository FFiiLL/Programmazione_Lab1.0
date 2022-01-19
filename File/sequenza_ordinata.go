package main

import(

    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

func main() {

    seq1, _ := os.Open(os.Args[1])
    seq2, _ := os.Open(os.Args[2])
    defer seq1.Close()
    defer seq2.Close()
    seqRes, _ := os.Create("sequenza_risultante.txt")
    defer seqRes.Close()
    
    scanner := bufio.NewScanner(seq1)
    var s1 string
    for scanner.Scan(){
       
        s1 += scanner.Text()
    }
    sequenza1 := strings.Split(s1," ")
    
    scanner = bufio.NewScanner(seq2)
    var s2 string
    for scanner.Scan(){
        s2 += scanner.Text()
    }
    sequenza2 := strings.Split(s2," ")
    
    var sequenzaInteri []int
    
    for _, valore := range sequenza1 {
        intero, _ := strconv.Atoi(valore)
        sequenzaInteri = append(sequenzaInteri, intero)
    }
    for _, valore := range sequenza2 {
        intero, _ := strconv.Atoi(valore)
        sequenzaInteri = append(sequenzaInteri, intero)
    }
    
    sort.Ints(sequenzaInteri)
    
    for _,intero := range sequenzaInteri{
        fmt.Fprintf(seqRes,"%d ",intero)
        
    }
    fmt.Fprintln(seqRes)
}
