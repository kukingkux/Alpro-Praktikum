package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luaslingkaran float64
	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)
	pi := math.Pi
	luaslingkaran = pi * r * r
	fmt.Printf("%.1f", luaslingkaran)
}
