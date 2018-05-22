package main

import "fmt"

func main() {
	// var nums []int
	// fmt.Println(nums, len(nums), cap(nums))

	// nums = make([]int, 5)
	// fmt.Println(nums, len(nums), cap(nums))

	// capitais := []string{"Lisboa"}
	// fmt.Println(capitais, len(capitais), cap(capitais))
	// capitais = append(capitais, "Brasília")
	// fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)
	cidades[0] = "São Francisco"
	cidades[1] = "Vancouver"
	cidades[2] = "São Paulo"
	cidades[3] = "Dublin"
	cidades[4] = "Londres"
	// fmt.Println(cidades, len(cidades), cap(cidades))

	// capitais[1] = "Brasilia"
	// fmt.Println(capitais, len(capitais), cap(capitais))

	// capitais = append(capitais, "Maputo")
	// fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}

	// Primeiro item começa com indice 0
	// Segundo item começa com indice 1

	capitaisEuropa := cidades[3:5]
	fmt.Println(capitaisEuropa)

	tempo1 := cidades[:2]
	fmt.Println(tempo1)

	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)

	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1])
	copy(cidades, temp)
	fmt.Println(temp)

}
