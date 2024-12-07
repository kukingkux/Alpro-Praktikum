package main

import "fmt"

func main() {
	var x, next, num, curr, total, i int64
	next = 1
	fmt.Scan(&x)
	for done := true; done; {
		if next <= x {
			num = (x / next) % 10
		} else {
			break
		}
		next = next * 10
		curr = (1 + num) / (1 + num)
		total += curr
		i++
		done = i <= total
	}
	fmt.Println(total)
}
