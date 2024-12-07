package main

import "fmt"

func main() {
	var x, next, curr int
	fmt.Scan(&x)
	next = 1
	for next < x {
		curr = (x / next) % 10
		next *= 10
		fmt.Println(curr)
		// if curr == 0 {
		// 	break
		// }
	}
}
