package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n            int
		r, t, volume float64
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r, &t)
		volume = 1.0 / 3 * math.Pi * r * r * t
		fmt.Println("Output:", volume)
	}
}
