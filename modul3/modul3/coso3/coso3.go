package main

import "fmt"

func main() {
	var rupiah, dollar float64
	fmt.Print("Masukkan nominal rupiah : ")
	fmt.Scan(&rupiah)
	dollar = (rupiah / 15000)
	fmt.Print("Jadi ", rupiah, " rupiah = ", dollar, " dollar")
}
