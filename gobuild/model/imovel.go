package model

import (
	"errors"
)

// Imovel armazena os valores
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// ErrValorDeImovelInvalido é apresentado quando é atribuido a um imóvel com valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

// ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuido um valor muito alto (acima da média do mercado)
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

// SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor

	return
}

// GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
