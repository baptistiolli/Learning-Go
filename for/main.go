package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Valor de i é", i)
	}

	valor := 0
	teste := true
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("Numero: ", valor)
	}

	for {
		valor--
		fmt.Println("Numero é:", valor)
		if valor == 0 {
			break
		}
	}

	texto := "I love writing using Go"
	for indice, letra := range texto {
		fmt.Printf("Texto:[%d] = %q\r\n", indice, letra)
	}
}
