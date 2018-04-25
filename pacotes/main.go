package main

import (
	"fmt"

	"github.com/baptistiolli/Learning-Go/variaveis/pacotes/operadora"
	"github.com/baptistiolli/Learning-Go/variaveis/pacotes/prefixo"
)

// Nome do usuário do sistema
var nomeDoUsuario = "Jhonny"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", nomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de TesteComPrefixo: %s\r\n", prefixo.TesteComPrefixo)
}
