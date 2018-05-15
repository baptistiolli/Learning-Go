package main

import (
	"fmt"

	"github.com/baptistiolli/Learning-Go/mapas/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apartamento"
	apto.X = 8
	apto.Y = 9
	apto.SetValor(150000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache/map? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache/map? ", len(cache))

}
