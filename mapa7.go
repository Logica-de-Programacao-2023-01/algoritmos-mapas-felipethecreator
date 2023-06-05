
import (
	"fmt"
	"strings"
)

func contarLetrasEPalavras(s string) map[string]map[string]int {

	mF := make(map[string]map[string]int)

	palavras := strings.Fields(s)

	chars := []string{}

	for _, ranP := range palavras {

		chars = strings.Split(ranP, "")

		mF[ranP] = make(map[string]int)

		for _, ranC := range chars {

			mF[ranP][ranC]++

		}

	}

	return mF

}

func main() {

	fmt.Println(contarLetrasEPalavras("caguei nas carça"))

}
