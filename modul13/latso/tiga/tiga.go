package main

import "fmt"

func main() {
	var target, perdonasi, donatur, akumulasi int
	fmt.Scan(&target)
	for done := false; !done; {
		fmt.Scan(&perdonasi)
		akumulasi += perdonasi
		done = akumulasi >= target
		donatur++
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, perdonasi, akumulasi)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", akumulasi, donatur)

}
