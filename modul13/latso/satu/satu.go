package main

import "fmt"

func main() {
	var x, next, i int64
	next = 1
	fmt.Scan(&x)
	for process := true; process; {
		next = next * 10
		process = next < x
		i = i + 1
	}
	fmt.Println(i)
}
