package agendador

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/4linux/go-labs/fs"
)

const (
	campoSegundo = 0
	campoMinuto  = 1
	campoHora    = 2
	ALL          = 999
)

func processarCampo(s string) int {
	valorStr := strings.TrimSpace(s)
	if valorStr == "*" {
		return ALL
	}

	valor, err := strconv.ParseInt(valorStr, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARN: Erro ao tentar converter: %s\n", valorStr)
		return -1
	}

	return int(valor)
}

func LerConfigAgendador(arquivo string) (tarefas []Tarefa, err error) {
	linhas, err := fs.LerArquivo(arquivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir arquivo '%s': %v\n", arquivo, err)
		return tarefas, errors.New("erro ao abrir arquivo")
	}

	for _, linha := range linhas {

		campos := strings.Split(linha, " ")
		if len(campos) < 3 {
			fmt.Fprintf(os.Stderr, "WARN: pulando a linha, erro interpretar a linha: %s\n", linha)
			continue
		}

		tarefa := Tarefa{
			Segundo: processarCampo(campos[campoSegundo]),
			Minuto:  processarCampo(campos[campoMinuto]),
			Hora:    processarCampo(campos[campoHora]),
			// tudo que sobrar, vira comando
			Command: strings.Join(campos[3:], " "),
		}
		tarefas = append(tarefas, tarefa)
	}

	return tarefas, nil
}
