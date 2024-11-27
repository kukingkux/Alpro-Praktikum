package main

import "fmt"

func main() {
	var umur int
	var kk bool
	fmt.Scan(&umur, &kk)
	if umur >= 17 && kk {
		fmt.Println("Bisa buat ktp")
	} else {
		fmt.Println("Belum bisa buat ktp")
	}
}
