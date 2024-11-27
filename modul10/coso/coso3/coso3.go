package main

import "fmt"

func main() {
	var x, d1, d2, d3, d4 int
	var teks string
	fmt.Print("Bilangan : ")
	fmt.Scan(&x)
	d4 = x % 10
	d3 = (x % 100) / 10
	d2 = (x % 1000) / 100
	d1 = x / 1000

	if d1 < d2 && d2 < d3 && d3 < d4 {
		teks = "terurut membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		teks = "terurut mengecil"
	} else {
		teks = "tidak terurut"
	}

	fmt.Println("Digit pada bilangan", x, teks)
}
