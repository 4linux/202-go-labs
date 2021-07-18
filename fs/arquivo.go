package fs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Abre um arquivo de sistema, retornando um ponteiro ao arquivo aberto.
// Antes e abrir o arquivo, o a função valida se o arquivo existe ou nao
// através da função `os.Stat()`, retornando um erro caso o arquivo nao
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
		return nil, fmt.Errorf("arquivo %s nao existe", caminhoArquivo)
	}

	arquivo, err = os.Open(caminhoArquivo)
	if err != nil {
		return nil, err
	}

	return arquivo, nil
}

// Fecha um arquivo aberto no sistema. Esta função e protegida
// caso o arquivo passado for `nulo` (nil).
func FecharArquivo(arquivo *os.File) {
	if arquivo != nil {
		arquivo.Close()
	}
}

// Realiza a leitura de um arquivo
func LerArquivo(caminhoArquivo string) (linhas []string, err error) {
	arquivo, err := AbrirArquivo(caminhoArquivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir arquivo '%s': %v\n", caminhoArquivo, err)
		return linhas, errors.New("erro ao abrir arquivo")
	}

	defer FecharArquivo(arquivo)

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	return linhas, nil
}
