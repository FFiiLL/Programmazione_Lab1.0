package main

import(

    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

type Prodotto struct{
    
    nome, codice string

}

type Magazzino map[Prodotto]int

func LeggiTesto() (testo string){

    scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){  
    testo += scanner.Text() + "\n"
  }  
    return testo[:len(testo)-1]

}

func nuovoMagazzino() (magazzino Magazzino){
    magazzino=make(map[Prodotto]int)
    return magazzino

}

func modificaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool){
        mIniziale := m
        m[p] += variazione
        
        if m[p] == 0{
            delete(m,p)
            return m,true
        }else if m[p] > 0{
            
            return m,true
        }else{
            return mIniziale,false
        }

}

func main() {

    testo := LeggiTesto()
    
    slComandi := strings.Split(testo,"\n")
    
    
    magazzino := nuovoMagazzino()
    var valido bool
    for i,prodotto := range slComandi{
        var pr Prodotto
        prodottoScomposto := strings.Split(prodotto,";")
        pr.nome = prodottoScomposto[0]
        pr.codice = prodottoScomposto[1]
        numero,_ := strconv.Atoi(prodottoScomposto[2] )
        magazzino,valido = modificaProdotto(magazzino,pr,numero)
        
        if !valido{

            fmt.Printf("Errore alla riga %d: Prodotto {%s %s} - disponibilità %d, richiesta %d",i+1,pr.codice,pr.nome,magazzino[pr],numero)
            return
        }
    
    }
    
   
    for k,v := range magazzino{
        fmt.Printf("Codice: %s, Prodotto: %s, Quantità: %d\n",k.codice,k.nome,v)
    }

    
    
}
