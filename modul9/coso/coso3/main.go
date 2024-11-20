package main

import "fmt"

func main() {
	var a int
	var bool bool
	fmt.Scan(&a)
	if a < 0 && a%2 == 0 {
		bool = true
	}
	fmt.Print(bool)
}
