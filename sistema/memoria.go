package sistema

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/4linux/go-labs/fs"
)

const (
	MEMINFO_FILE = "/proc/meminfo"

	// nome dos campos do arquivo /proc/meminfo
	MEMINFO_TOTAL     = "MemTotal"
	MEMINFO_FREE      = "MemFree"
	MEMINFO_AVAILABLE = "MemAvailable"
	MEMINFO_CACHED    = "Cached"
	MEMINFO_BUFFERED  = "Buffers"
)

type InfoMemoria struct {
	Total    int64
	Cached   int64
	Buffered int64
	Free     int64
}

func extrairLinha(s string) (campo string, valor int64) {
	var err error

	linha := strings.Split(s, ":")
	if len(linha) != 2 {
		// vamos ignorar erros, segue o jogo!
		fmt.Println("1")
		return "", -1
	}

	linhaValor := strings.Split(strings.Trim(linha[1], " "), " ")
	if len(linhaValor) != 2 {
		// vamos ignorar erros, segue o jogo!
		return "", -1
	}

	valor, err = strconv.ParseInt(strings.Trim(linhaValor[0], " "), 10, 64)
	if err != nil {
		// vamos ignorar erros, segue o jogo!
		return "", -1
	}

	campo = strings.Trim(linha[0], " ")
	return campo, valor
}

func GetInfoMemoria(caminhoArquivo string) (info *InfoMemoria) {
	info = new(InfoMemoria)

	arquivo, err := fs.AbrirArquivo(caminhoArquivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir arquivo '%s': %v\n", MEMINFO_FILE, err)
		return nil
	}
	defer fs.FecharArquivo(arquivo)

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		campo, valor := extrairLinha(scanner.Text())
		switch campo {
		case MEMINFO_TOTAL:
			info.Total = valor
		case MEMINFO_FREE:
			info.Free = valor
		case MEMINFO_CACHED:
			info.Cached = valor
		case MEMINFO_BUFFERED:
			info.Buffered = valor
		}
	}

	return info
}

func GetInfoMemoriaDefault() (info *InfoMemoria) {
	return GetInfoMemoria(MEMINFO_FILE)
}

// Returna a utilização de memoria. O valor e computado subtraindo os valores
// used + buffered + used do valor total.
func (m *InfoMemoria) UsoPerct() float64 {
	uso := m.Free + m.Buffered + m.Cached
	util := float64(uso) / float64(m.Total) * 100

	return 100 - util
}

// Retorna a utilização, em porcentagem, da utilização da memória. O valor é calculado
// total - free - buffers - cache
func MemUso() float64 {
	info := GetInfoMemoriaDefault()
	if info == nil {
		return -1
	}
	return info.UsoPerct()
}
