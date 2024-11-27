package main

import "fmt"

func main() {
	var bil, factor int
	var prime bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	fmt.Print("Faktor: ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Printf("%v ", i)
			factor++
		}
	}

	if factor == 2 {
		prime = true
	}
	fmt.Printf("\nPrima: %v", prime)
}
