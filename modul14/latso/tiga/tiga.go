package main

import "fmt"

func main() {
	const correctWarna1 = "merah"
	const correctWarna2 = "kuning"
	const correctWarna3 = "hijau"
	const correctWarna4 = "ungu"

	var (
		warna1, warna2, warna3, warna4 string
		total                          int
		isTrue                         bool
	)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna1)
		fmt.Scan(&warna2)
		fmt.Scan(&warna3)
		fmt.Scan(&warna4)
		if warna1 == correctWarna1 && warna2 == correctWarna2 && warna3 == correctWarna3 && warna4 == correctWarna4 {
			total = total + 1
		}
	}
	isTrue = total == 5

	fmt.Println("BERHASIL: ", isTrue)
}
