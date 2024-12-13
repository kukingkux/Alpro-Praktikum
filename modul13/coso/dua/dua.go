package main

import "fmt"

func main() {
	var (
		x int
	)
	for cout := true; cout; {
		fmt.Scan(&x)
		cout = x <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", x)
}
