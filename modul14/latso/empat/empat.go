package main

import "fmt"

func main() {
	var input, allInput string
	var n int
	for {
		fmt.Printf("Bunga %d: ", n+1)
		fmt.Scan(&input)
		if input == "SELESAI" {
			break
		}
		allInput += input + " - "
		n++
	}
	fmt.Println("Pita:", allInput)
	fmt.Println("Bunga:", n)
}
