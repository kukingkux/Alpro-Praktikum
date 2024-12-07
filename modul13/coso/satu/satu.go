package main

import "fmt"

func main() {
	var (
		kata string
		x    int
	)

	fmt.Scan(&kata, &x)
	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= x)
	}
}
