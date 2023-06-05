package main

import (
	"fmt"
	"strings"
)

func contarPalavras2(s string) map[string]int {

	m1 := make(map[string]int)
	conj := strings.Fields(s)

	for _, ranConj := range conj {

		m1[ranConj] += 1

	}

	return m1

}

func unirMapas2(mapas []map[string]int) map[string]int {

	bigMap := make(map[string]int)

	for _, ranMapas := range mapas {

		for chave, valor := range ranMapas {

			bigMap[chave] += valor

		}

	}

	return bigMap

}

func main() {

	a := "muçilon"
	b := "adulto"
	c := "velho"

	conj := []map[string]int{contarPalavras2(a), contarPalavras2(b), contarPalavras2(c)}

	fmt.Println(unirMapas2(conj))

}
