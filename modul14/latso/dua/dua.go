package main

import "fmt"

func main() {
	var bil, factor int
	var text string

	fmt.Scan(&bil)
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			factor = factor + 1
		}
	}
	if factor == 2 {
		text = "prima"
	} else {
		text = "bukan Prima"
	}
	fmt.Print(text)
}
