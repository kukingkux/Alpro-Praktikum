package main

import "fmt"

func main() {
	var bil, jumlah, rerata, i float64
	for {
		fmt.Scan(&bil)
		if bil == 9999 {
			break
		}
		i = i + 1
		jumlah = jumlah + bil
	}
	rerata = jumlah / i
	fmt.Println(rerata)
}
