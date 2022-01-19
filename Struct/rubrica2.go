package main

import(

    "fmt"

)
type Indirizzo struct{
    via, CAP, citta string
    numeroCivico uint
}

type Contatto struct{
    cognome, nome, telefono string
    indirizzo Indirizzo
}

type Rubrica map[string]Contatto

func NuovaRubrica()  Rubrica{
    
    return make(Rubrica)
}

func InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, CAP, citta string, telefono string) Rubrica {
    chiave := cognome + nome
    for k := range r{
        if k == chiave{
            return r
        }
    }
    
    r[chiave] = Contatto{cognome,nome,telefono,Indirizzo{via,CAP,citta,numero}}
    return r
}

func EliminaContatto(r Rubrica, cognome, nome string) Rubrica {
    chiave := cognome + nome
    for k := range r{
        if k == chiave{
            delete(r,chiave)
        }
    }
    
    return r
    
}

func StampaContatto(c Contatto){
    fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n", c.nome, c.cognome, c.indirizzo.via, c.indirizzo.numeroCivico, c.indirizzo.CAP, c.indirizzo.citta, c.telefono)
}

func StampaRubrica(r Rubrica){
    contatore := 1
    fmt.Println("-----RUBRICA----")
    for _,contatto := range r{
        fmt.Print("[",contatore,"] - ")
        StampaContatto(contatto)
        
        contatore++
    }
}

func AggiornaContatto(r Rubrica, cognome, nome string, via string, numero uint, CAP, citta string, telefono string) Rubrica {
    
    chiave := cognome + nome
    for k,_ := range r{
    
        if k == chiave{
            r[k] = Contatto{cognome,nome,telefono,Indirizzo{via,CAP,citta,numero}}
            return r
        }
    
    }
    return r
}

func main() {

    r := NuovaRubrica()
   
    r = InserisciContatto(r,"Rossi","Mario","Via Celoria",18,"20122","Milano","02503111")
    r = InserisciContatto(r,"Rossi","Anna","Via Festa del Perdono",11,"20122","Milano","02503111")
    r = InserisciContatto(r,"Rossi","Carlo","Via Festa del Perdono",11,"20122","Milano","02503111")
    StampaRubrica(r)
    r = EliminaContatto(r,"Rossi","Mario")
    StampaRubrica(r)
    r = AggiornaContatto(r,"Rossi","Anna","Via Santa Sofia",24,"20122","Milano","")
    StampaRubrica(r)
    
}
