package main

import "fmt"

func main() {
	var x, fx float64
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fx = (2.0 / (x + 5.0)) + 5
	fmt.Print(fx)
}
