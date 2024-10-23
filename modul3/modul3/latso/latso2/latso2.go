package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		r            int
		volume, luas float64
	)
	fmt.Print("Masukkan jari-jari bola : ")
	fmt.Scanln(&r)
	rFloat := float64(r)
	volume = (4.0 / 3.0) * math.Pi * (rFloat * rFloat * rFloat)
	luas = 4 * math.Pi * (rFloat * rFloat)
	fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}
