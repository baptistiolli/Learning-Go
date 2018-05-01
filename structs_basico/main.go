package main

import "fmt"

// Imovel é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu ape", 760000}
	fmt.Printf("O apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacará",
		X:     22,
		valor: 55,
	}
	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Home sweet home"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 32

	fmt.Printf("A casa é: %+v\r\n", casa)
}
