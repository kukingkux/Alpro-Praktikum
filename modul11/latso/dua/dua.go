package main

import "fmt"

func main() {
	var (
		kendaraan, waktu string
		durasi, biaya    int
	)

	fmt.Scanf("%s %d %s", &kendaraan, &durasi, &waktu)
	if waktu == "menit" && durasi <= 60 {
		durasi = 1
	} else if waktu == "menit" && durasi > 60 {
		durasi = durasi/60 + 1
	}

	switch {
	case (kendaraan == "mobil" || kendaraan == "Mobil") && (waktu == "jam" || waktu == "menit"):
		for i := 0; i < durasi; i++ {
			biaya += 5000
		}
		fmt.Printf("Rp %d", biaya)
	case (kendaraan == "motor" || kendaraan == "Motor") && (waktu == "jam" || waktu == "menit"):
		for i := 0; i < durasi; i++ {
			biaya += 2000
		}
		fmt.Printf("Rp %d", biaya)
	case (kendaraan == "truk" || kendaraan == "Truk") && (waktu == "jam" || waktu == "menit"):
		for i := 0; i < durasi; i++ {
			biaya += 8000
		}
		fmt.Printf("Rp %d", biaya)
	default:
		fmt.Print("Invalid Input")
	}
}
