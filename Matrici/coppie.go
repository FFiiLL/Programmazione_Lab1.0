package main

import(

    "os"
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	return s
}

func AssegnaSliceBidimensionale(s [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}
}

func Coppie (s [][]int) (coppie [][]int){
    for i, riga := range s{
        for j, valore := range riga{
            if valore == 1{
                coppie = append(coppie, []int{i,j})
            }
        }
        
    }
    return 
}

func main() {

    n, _ := strconv.Atoi(os.Args[1])
    
    slice := CreaSliceBidimensionale(n)
    AssegnaSliceBidimensionale(slice)
    
    for _, riga := range slice{
        for _, valore := range riga{
            fmt.Print(valore," ")
        }
        fmt.Println()
    }
    
    coppie := Coppie(slice)
    
    for _, coppia := range coppie{
        fmt.Println(coppia)
    }

}
