package main

import "fmt"

func main() {
	var (
		pass string
	)
	fmt.Scan(&pass)
	for pass != "12345abcde" {
		fmt.Scan(&pass)
	}
	fmt.Println("Berhsil Login")
}
