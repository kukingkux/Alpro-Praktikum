package main

import "fmt"

func main() {
	var (
		hasil, j int
	)
	fmt.Scan(&j)
	for i := 1; i <= j; i++ {
		hasil = hasil + i
	}
	fmt.Print(hasil)
}
