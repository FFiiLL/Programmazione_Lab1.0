package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFiltro(t *testing.T) {

	RunTest(t, []string{"MANDARINO"}, "", `    O    
    N    
    I    
    R    
ONIRADNAM
    D    
    N    
    A    
    M    
`)

	RunTest(t, []string{"bignè"}, "", `  è  
  n  
èngib
  i  
  b  
`)

	RunTest(t, []string{"ƤƦŌĢɭăƁ"}, "", `   Ɓ   
   ă   
   ɭ   
ƁăɭĢŌƦƤ
   Ō   
   Ʀ   
   Ƥ   
`)

}

func RunTest(t *testing.T, args []string, stdin string, stdout string) { //si possono fare slice per stdin e stdout ma attento ai \n finali, se si vuole testare anche il tempo bisognerebbe fare controlli aggiuntivi

	os.Args = append([]string{""}, args...)

	readStdin, writeStdin, _ := os.Pipe()     //pipe written to writeStdin can be read from readStdin
	readStdin, os.Stdin = os.Stdin, readStdin //now we can write program input to writeStdin

	readStdout, writeStdout, _ := os.Pipe()         //pipe written to writeStdout can be read from readStdout
	writeStdout, os.Stdout = os.Stdout, writeStdout //now we can read program output from readStdout

	fmt.Fprint(writeStdin, stdin)

	main()

	//revert correct stdin and stdout files
	readStdin, os.Stdin = os.Stdin, readStdin
	writeStdout, os.Stdout = os.Stdout, writeStdout
	readStdin.Close()
	writeStdin.Close()
	writeStdout.Close()

	result, _ := ioutil.ReadAll(readStdout)

	readStdout.Close()

	if stdout != string(result) {
		t.Error(fmt.Sprintf("\nArgs: %s\nInput: %s\nExpected result:\n-------\n%s\n-------\nGiven result:\n-------\n%s\n-------\n", strings.Join(args, " "), stdin, stdout, string(result)))
	}
}


 
