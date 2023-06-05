package main

import "fmt"

// Escreva uma função que receba dois mapas
// e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
// Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.

func unirMapas(m1 map[string]int, m2 map[string]int) map[string]int {

	m3 := make(map[string]int)

	for range1, range2 := range m1 {

		m3[range1] = range2

	}

	for range3, range4 := range m2 {

		m3[range3] = range4

	}

	return m3

}

func main() {

	m1 := map[string]int{

		"macaco":       1,
		"aranha": 4,
		"kiko":       2,
	}

	m2 := map[string]int{

		"cavalo": 2,
		"rato":              9999,
		"1":                 2,
	}

	fmt.Println(unirMapas(m1, m2))

}
