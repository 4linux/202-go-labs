package numeros

import (
	"math/rand"
)

func GeradorLista(n int) []int {
	lista := make([]int, n)

	for i := 0; i < n; i++ {
		lista[i] = rand.Int() % 100
	}

	return lista
}
