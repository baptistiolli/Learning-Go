package main

import (
	"encoding/json"
	"fmt"

	"github.com/baptistiolli/Learning-Go/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	if err := casa.SetValor(11000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor, revise a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))

}
