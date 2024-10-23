package main

import "fmt"

func main() {
	var (
		n, j, pow int
	)
	pow = 1
	fmt.Scan(&n, &j)
	for i := 1; i <= j; i++ {
		pow = pow * n
	}
	fmt.Print(pow)
}
