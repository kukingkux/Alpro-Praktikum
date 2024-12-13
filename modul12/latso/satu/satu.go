package main

import "fmt"

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"
	var i int
	var username, password string
	var isPass bool

	fmt.Scan(&username, &password)
	isPass = username == correctUsername && password == correctPassword
	for isPass == false {
		fmt.Scan(&username, &password)
		isPass = username == correctUsername && password == correctPassword
		i++
	}

	fmt.Printf("%d percobaan gagal login", i)
}
