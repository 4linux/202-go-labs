package fs

import (
	"testing"
)

func TestAbrirArquivoNaoExiste(t *testing.T) {
	arquivo, err := AbrirArquivo("/tmp/nao_existente.1234aa")
	if arquivo != nil && err == nil {
		t.Errorf("abrirArquivo() com arquivo inexistente deveria retorna erro")
	}
}

func TestAbrirArquivoMemoria(t *testing.T) {
	if _, err := AbrirArquivo("/proc/meminfo"); err != nil {
		t.Errorf("abrirArquivo() com arquivo inexistente deveria retorna erro")
	}
}

func TestFecharArquivoNulo(t *testing.T) {
	arquivo, err := AbrirArquivo("/tmp/nao_existente.1234aa")
	if err == nil {
		t.Errorf("abrirArquivo() com arquivo inexistente ndeveria retornar erro")
	}
	FecharArquivo(arquivo)
}
