package main

import (
	"fmt"
)

func main() {

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1
	fmt.Println("Qual a capacidade do array?", len(teste))

	cantores := [2]string{"Daniel", "Beatles"}
	fmt.Printf("Conteudo do array \n\r%v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Sao Paulo", "Luanda", "Maputo"}

	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\n\r", indice, cidade)
	}

}
