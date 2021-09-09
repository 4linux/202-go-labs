package main

import (
	"fmt"
	"os"

	"github.com/raffs/go-labs/sistema"
)

func main() {
	memUso := sistema.MemUso()
	if memUso < 0 {
		fmt.Fprintln(os.Stderr, "Nao foi possivel coletar uso de memoria")
		os.Exit(1)
	}
	fmt.Println("Uso memoria(%):", memUso)
}
