package main

import "fmt"

func main() {
	var x, j, temp, s1, s2 int
	fmt.Scan(&x)
	s1 = 0
	s2 = 1
	for j < x {
		fmt.Print(s1, " ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
		j++
	}
}
