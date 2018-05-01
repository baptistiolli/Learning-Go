package model

// Armazena os dados de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y`
	Nome  string `json:"nome"`
	valor int
}

// Define o valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

// Retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
