package matematica

/*
Executa quaisquer tipos de cálculos, basta enviar a funcao desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numberoB int) int {
	return funcao(numeroA, numberoB)
}

/*
Multiplica dois números
*/
func Multiplicador(x int, y int) int {
	return x * y
}

// Efetua a divisão de dois numeros
func Divisor(numeroA int, numberoB int) (resultado int) {
	resultado = numeroA / numberoB
	return
}

// Divisor com retorno do resto da divisao
func DivisorComResto(numeroA int, numberoB int) (resultado int, resto int) {
	resultado = numeroA / numberoB
	resto = numeroA % numberoB
	return
}
