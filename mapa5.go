package main

import (
	"fmt"
	"strings"
)

func frequenciaLetras(s string) map[string]int {

	m1 := make(map[string]int)
	chars := strings.Split(s, "")

	for _, ranC := range chars {

		m1[ranC] += 1

	}

	return m1

}

func main() {

	fmt.Println(frequenciaLetras("caguei nas calça"))

}
