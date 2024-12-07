package main

import (
	"fmt"
)

func main() {
	var x float64
	var itg int
	fmt.Scan(&x)
	itg = int(x)
	for done := false; !done; {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)
		done = x >= 1.0+float64(itg)
	}
}
