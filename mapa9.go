package main

import "fmt"

func fibonacci(n int) map[int]int {

	fibonas := make(map[int]int)

	for i := 0; len(fibonas) != n; i++ {

		if i < 2 {

			fibonas[i] = i

		} else {

			fibonas[i] = fibonas[i-2] + fibonas[i-1]

		}

	}

	return fibonas

}

func main() {

	fmt.Println(fibonacci(20))

}
