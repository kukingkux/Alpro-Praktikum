package main

import "fmt"

func main() {
	var x, y int
	var factorXY, factorYX bool
	fmt.Scan(&x, &y)
	if y%x == 0 {
		factorYX = true
	}
	if x%y == 0 {
		factorXY = true
	}
	fmt.Println(factorYX, factorXY)
}
