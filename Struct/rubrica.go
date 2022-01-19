package main

import(

    "fmt"
    "os"
    "bufio"
    "strings"
    
    "strconv"
)

type Indirizzo struct{
    
    via string
    CAP string
    citta string
    numeroCivico uint
  
}

type Contatto struct{
    
    cognome string
    nome string
    telefono string
    indirizzo Indirizzo
    
}

type Rubrica []Contatto

func NuovaRubrica() (rubrica Rubrica){
    return
}

func StampaContatto(c Contatto){
    fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n",c.nome,c.cognome,c.indirizzo.via,c.indirizzo.numeroCivico,c.indirizzo.citta,c.indirizzo.CAP,c.telefono)
}

func StampaRubrica(r Rubrica){
    fmt.Println("rubrica")
    for i, contatto := range r{
        fmt.Print("[",i+1,"] - ") 
        StampaContatto(contatto)
        
    }
}

func InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, CAP, citta string, telefono string) Rubrica{
    flag := true
    
    for _, contatto := range r{
        
        if contatto.nome == nome && contatto.cognome == cognome{
            flag = false
        }
    
    }
        var ind Indirizzo 
        ind.via = via
        ind.CAP = CAP
        ind.citta = citta
        ind.numeroCivico = numero
        var cont Contatto 
        cont.cognome = cognome
        cont.nome = nome
        cont.telefono = telefono
        cont.indirizzo = ind
        
    if flag{
       
        
        r = append(r,cont)
        return r
    }else{
        return r
    }

}
func AggiornaContatto(r Rubrica, cognome, nome string, via string, numero uint, CAP, citta string, telefono string) Rubrica{
    
    
    for i, contatto := range r{
        
        if contatto.nome == nome && contatto.cognome == cognome{
            var ind Indirizzo 
            ind.via = via
            ind.CAP = CAP
            ind.citta = citta
            ind.numeroCivico = numero
            var cont Contatto 
            cont.cognome = cognome
            cont.nome = nome
            cont.telefono = telefono
            cont.indirizzo = ind
            r[i] = cont
            return r
        }
    
    }
    
    return r
        
        
    

}


func EliminaContatto(r Rubrica, cognome, nome string) Rubrica{
   
   for i, contatto := range r{
        
        if cognome == contatto.cognome && nome == contatto.nome{
            r = append(r[:i],r[i+1:]...)
        }
        
   } 
   return r
}

func LeggiTesto() (testo string){
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan(){
        
        testo += scanner.Text() + "\n"
        
    }
    return
} 

func main() {

    rubrica := NuovaRubrica()
    testo := LeggiTesto()
    
    comandi := strings.Split(testo,"\n")
   
    
    for _, comando := range comandi{
        comandoSeparato := strings.Split(comando,";")
        switch comandoSeparato[0]{
            
            case "I":
            nCivico,_ := strconv.Atoi(comandoSeparato[4])
            rubrica = InserisciContatto(rubrica, comandoSeparato[1], comandoSeparato[2], comandoSeparato[3], uint(nCivico), comandoSeparato[5], comandoSeparato[6],comandoSeparato[7])
            break
            case "E":
            rubrica = EliminaContatto(rubrica,comandoSeparato[1],comandoSeparato[2])
            break
            case "S":
            StampaRubrica(rubrica)
            break
            case "A":
            nCivico,_ := strconv.Atoi(comandoSeparato[4])
            rubrica = AggiornaContatto(rubrica, comandoSeparato[1], comandoSeparato[2], comandoSeparato[3], uint(nCivico), comandoSeparato[5], comandoSeparato[6],comandoSeparato[7])
            break
            default:
                break
        }
    }
    
    
   
}
