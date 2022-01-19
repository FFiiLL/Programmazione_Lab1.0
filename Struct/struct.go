package main
import(
    "fmt"
)


type Coordinate2D struct{
    
        x float64
        y float64   
        
}


func cambiaPunto(p *Coordinate2D){

    p.x = 0
    p.y = 0

}


func main() {

    var punto Coordinate2D
    
    punto.x = 2.3
    punto.y = 4.5
    
    cambiaPunto(&punto)
    
    fmt.Println(punto)    


}
