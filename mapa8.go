package main

import "fmt"

func igualarDespesas(m1 map[string]float64) {

	maior := 0.0

	for _, ranM1 := range m1 {

		if ranM1 > maior {

			maior = ranM1

		}

	}

	for nome, quantia := range m1 {

		quantia = maior - quantia

		fmt.Printf("%s deve pagar %.2f reais para  igualar.\n", nome, quantia)

	}

}

func main() {

	m1 := make(map[string]float64)
	m1["Repet"] = 10000
	m1["Jubis"] = 100
	m1["Melo"] = 500

	igualarDespesas(m1)

}
