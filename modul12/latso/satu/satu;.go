package main

import "fmt"

func main() {
	var i int
	var user, pass string
	fmt.Scan(&user, &pass)
	for user != "Admin" && pass != "Admin" {
		fmt.Scan(&user, &pass)
		i++
	}
	fmt.Printf("%d percobaan gagal login", i)
}
