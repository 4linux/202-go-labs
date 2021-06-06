package numeros

func DividirEm(lista []int, partes int) (subLista [][]int) {
	tamanho := len(lista)
	blocos := tamanho / partes

	for inicio := 0; inicio < tamanho; inicio += blocos {
		final := inicio + blocos

		if final > tamanho {
			final = tamanho
		}

		subLista = append(subLista, lista[inicio:final])
	}

	return subLista
}
