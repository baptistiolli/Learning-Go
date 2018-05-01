package main

import (
	"fmt"
)

func main() {
	meses := 0
	situacao := true
	cidade := "Osvaldo Cruz"

	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo!")
	}

	if situacao {
		fmt.Println("Credor está devendo!")
	}

	if !situacao {
		fmt.Println("Credor está limpo!")
	}

	if cidade == "Osvaldo Cruz" {
		fmt.Println("O cliente vive na cidade de", cidade)
	}

	if descricao, status := verificaTempoDeInadimplencia(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar!")

}

func verificaTempoDeInadimplencia(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo!"
		return
	}
	descricao = "O cliente está sem débitos!"
	return
}
