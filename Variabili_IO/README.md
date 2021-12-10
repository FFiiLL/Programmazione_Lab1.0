esercizi da https://github.com/UnimiDiProg1/Lab_2021-2022/tree/main/Laboratorio_2/Variabili_e_IO

```
package main

import("fmt";"math")

func main() {

    var n,l int
    
    fmt.Println("inserisci nnumero lati poligono =")
    fmt.Scan(&n)
    fmt.Println("inserisci lunghezza lati poligono =")
    fmt.Scan(&l)
    var area float64 = (float64(n) * math.Pow(float64(l), 2)) / (4 * math.Tan(math.Pi/float64(n)))
    fmt.Println("area: ",area)
}
```
