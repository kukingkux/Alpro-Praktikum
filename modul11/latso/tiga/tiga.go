package main

import "fmt"

func main() {
	var bilangan, b1, b2 int
	var kategori string
	fmt.Scan(&bilangan)

	switch {
	case bilangan%2 != 0 && bilangan == 5:
		b1 = bilangan
		b2 = bilangan + 1
		bilangan = b1 + b2
		kategori = "Ganjil"
		fmt.Printf("Kategori : Bilangan %s\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d", kategori, b1, b2, bilangan)

	case bilangan%2 == 0 && bilangan%10 != 0:
		b1 = bilangan
		b2 = bilangan + 1
		bilangan = b1 * b2
		kategori = "Genap"
		fmt.Printf("Kategori : Bilangan %s\nHasil perkalian dengan bilangan berikutnya %d * %d = %d", kategori, b1, b2, bilangan)

	case bilangan%5 == 0 && bilangan%10 != 0:
		b1 = bilangan
		bilangan = b1 * b1
		kategori = "Kelipatan 5"
		fmt.Printf("Kategori : Bilangan %s\nHasil kuadrat dari %d ^2 = %d", kategori, b1, bilangan)

	case bilangan%10 == 0:
		b1 = bilangan
		bilangan = b1 / 10
		kategori = "Kelipatan 10"
		fmt.Printf("Kategori : Bilangan %s\nHasil pembagian antara %d / 10 = %d", kategori, b1, bilangan)
	}
}
