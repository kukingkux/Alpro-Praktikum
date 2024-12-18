package main

import "fmt"

func main() {
	var bil, j int
	fmt.Scan(&bil)
	for j = 1; j <= bil; j++ {
		if j%2 != 0 {
			fmt.Print(j, " ")
		}
	}
}
