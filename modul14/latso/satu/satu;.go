package main

import "fmt"

func main() {
	var bil, total int
	fmt.Scan(&bil)
	for i := 1; i <= bil; i++ {
		if i%2 != 0 {
			total = total + 1
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", total)
}
