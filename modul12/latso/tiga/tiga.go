package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scan(&x, &y)
	for x >= y {
		x = x - y
		i = i + 1
	}
	fmt.Println(i)
}
