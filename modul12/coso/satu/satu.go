package main

import "fmt"

func main() {
	var (
		a int
	)
	fmt.Scan(&a)
	for a > 1 {
		fmt.Print(a, " x ")
		a = a - 1
	}
	fmt.Println(1)
}
