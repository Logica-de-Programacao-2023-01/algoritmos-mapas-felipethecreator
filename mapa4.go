package main

import (
	"fmt"
	"sort"
	"strings"
)

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {

	chars := strings.Split(s, "")

	sort.Strings(chars)

	return strings.Join(chars, "")
}

func main() {

	palavras := []string{"mono", "rato", "toma", "trab", "noom", "tora", "otma", "ratb"}

	fmt.Println(encontrarAnagramas(palavras))

}
