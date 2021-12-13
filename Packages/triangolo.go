package triangolo
import "math"
import "fmt"

type Triangolo struct{
    
     a float64
     b float64
     c float64
    

}
func AreaPeri(t Triangolo){
    
    p := (t.a + t.b + t.c) 
    fmt.Println("Perimetro:",p)
    p/=2
    fmt.Println("area:", math.Sqrt(p * (p-t.a) * (p-t.b) * (p-t.c)))


}
func NuovoTriangolo(l1,l2,l3 float64) Triangolo {

    var tri Triangolo 
    
    tri.a = l1
    tri.b = l2
    tri.c = l3
    
    return tri
}


