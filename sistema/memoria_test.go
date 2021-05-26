package sistema

import (
	"testing"
)

func testarMemUso(t *testing.T) {
	uso := MemUso()
	if uso < 0 {
		t.Errorf("MemUso() retornou um valor negavito (%f)", uso)
	}
}

func testarParseDeMemoria(t *testing.T) {
	info := GetInfoMemoriaDefault()
	if info.Total <= 0 {
		t.Errorf("info.Total retornou zero ou negativo (%d)", info.Total)
	}
	if info.Buffered <= 0 {
		t.Errorf("info.Buffered retornou zero ou negativo (%d)", info.Buffered)
	}
	if info.Cached <= 0 {
		t.Errorf("info.Cached retornou zero ou negativo (%d)", info.Total)
	}
	if info.Free <= 0 {
		t.Errorf("info.Free retornou zero ou negativo (%d)", info.Total)
	}
}

func testarUsoPerct(t *testing.T) {
	info := InfoMemoria{Total: 10, Free: 1, Buffered: 1, Cached: 1}
	perc := info.UsoPerct()
	if perc != 70 {
		t.Errorf("info.UsoPerct() retornou %f mas deveria retornar 70", perc)
	}
}
