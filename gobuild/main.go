package main

import (
	"encoding/json"
	"fmt"

	"github.com/baptistiolli/Learning-Go/gobuild/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa do Jhonny"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100)

	fmt.Printf("A casa é: %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
