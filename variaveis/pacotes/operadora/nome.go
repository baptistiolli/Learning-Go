package operadora

import (
	"strconv"

	"github.com/baptistiolli/Learning-Go/variaveis/pacotes/prefixo"
)

// Nome da operadora que esta sendo utilizada
var NomeOperadora = "XPTO Telecom"

// Prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
