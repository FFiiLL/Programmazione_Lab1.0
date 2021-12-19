package main

import(

    "fmt"
    "math"

)

func Distanza (x1, y1, x2, y2 float64) float64 {
    
    return math.Sqrt((math.Pow(x1-x2,2)+math.Pow(y1-y2,2)))
    
    
}

func main() {

    var xA,yA,xB,yB,xC,yC float64
    fmt.Print("Inserisci coordinate xy punto A: ")
    fmt.Scan(&xA,&yA)
    fmt.Print("Inserisci coordinate xy punto B: ")
    fmt.Scan(&xB,&yB)
    fmt.Print("Inserisci coordinate xy punto C: ")
    fmt.Scan(&xC,&yC)
    
    AB := Distanza(xA,yA,xB,yB)
    BC := Distanza(xB,yB,xC,yC)
    CA := Distanza(xC,yC,xA,yA)
    
    if AB == BC && BC == CA{
        fmt.Println("triangolo EQUIlatero lato",AB)
    }else if AB == BC && BC != CA{
        fmt.Println("ISOSCELE AB BC uguali lunghi:",AB,"CA:",CA)
    }else if AB == CA && BC != AB{
        fmt.Println("ISOSCELE AB CA uguali lunghi:",AB,"BC:",BC)
    }else if CA == BC && CA != AB{
        fmt.Println("ISOSCELE CA BC uguali lunghi:",CA,"AB:",AB)
    }else{
        fmt.Println("SCALENO LATI:",AB,BC,CA)
    }

}
