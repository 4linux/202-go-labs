package sistema

import (
	"testing"
)

func testarCpuUso(t *testing.T) {
	uso := CpuUso()
	if uso < 0 {
		t.Errorf("CpuUso() retornou um valor negativo (%f)", uso)
	}
}
