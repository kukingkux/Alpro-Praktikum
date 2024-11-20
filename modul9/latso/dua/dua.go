package main

import "fmt"

func main() {
	var n int
	var cek string
	fmt.Scan(&n)
	cek = "bukan"
	if n < 0 && n%2 == 0 {
		cek = "genap negatif"
	}
	fmt.Print(cek)
}
