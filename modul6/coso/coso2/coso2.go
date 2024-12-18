package main

import "fmt"

func main() {
	var a, hasil string
	a = "* "
	for i := 0; i < 5; i++ {
		hasil = hasil + a
		fmt.Println(hasil)
	}
}
