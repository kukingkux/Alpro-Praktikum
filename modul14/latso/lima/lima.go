package main

import "fmt"

func main() {
	var (
		kiri, kanan, totalBerat, selisihBerat float64
		oleng                                 bool
	)

	for totalBerat <= 150 && (kiri >= 0 && kanan >= 0) {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri)
		fmt.Scan(&kanan)
		if kanan >= kiri {
			selisihBerat = kanan - kiri
		} else if kiri >= kanan {
			selisihBerat = kiri - kanan
		}

		totalBerat = kanan + kiri
		if totalBerat >= 150 {
			break
		}

		oleng = selisihBerat >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
	fmt.Println("Proses selesai.")
}
