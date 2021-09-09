package sistema

import (
	"math/rand"
	"time"
)

// GetInfoCpuDefault retorna uma número aleatorio, por enquanto.
func GetInfoCpuDefault() float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(100))
}

func CpuUso() float64 {
	uso := GetInfoCpuDefault()
	return uso
}
