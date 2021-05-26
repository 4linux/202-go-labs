//
package sistema

import (
	"fmt"
	"os"
)

// Abre um arquivo de sistema, retornando um pointeiro ao arquivo aberto.
// Antes e abrir o arquivo, o a funcao valida se o arquivo existe ou nao
// atraves da funcao `os.Stat()`, retornando um erro caso o arquivo nao
// exista.
//
//  Exemplo:
//
//     arquivo, err := AbrirArquivo("/caminho/para/o/arquivo")
//     if err != nil {
//         fmt.Println("Erro ao abrir o arquivo!")
//     }
//
func AbrirArquivo(caminhoArquivo string) (arquivo *os.File, err error) {
	// verificar se o arquivo existe
	if _, err = os.Stat(caminhoArquivo); err != nil {
		return nil, fmt.Errorf("Arquivo %s nao existe", caminhoArquivo)
	}

	arquivo, err = os.Open(caminhoArquivo)
	if err != nil {
		return nil, err
	}

	return arquivo, nil
}


// Fecha um arquivo aberto no sistema. Esta funcao e protegida
// caso o arquivo passado for `nulo` (nil).
func FecharArquivo(arquivo *os.File) {
	if arquivo != nil {
		arquivo.Close()
	}
}
