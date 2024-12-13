package main

import "fmt"

func main() {
	var x float64
	var itg int
	fmt.Scan(&x)
	itg = int(x) + 1
	for done := false; !done; {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)
		done = x > float64(itg)-0.11
	}
	fmt.Println(itg)
}
