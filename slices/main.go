package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "Londres"
	cidades[1] = "Vancouver"
	cidades[2] = "Dublin"
	cidades[3] = "São Francisco"
	fmt.Println(cidades, len(cidades), cap(cidades))

	capitais[1] = "Brasilia"
	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Maputo")
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}
}
