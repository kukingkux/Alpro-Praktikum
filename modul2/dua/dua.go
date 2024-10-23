package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		nama, kelas string
		nim         int64
	)
	namaScan := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama lengkap: ")
	if namaScan.Scan() {
		nama = namaScan.Text()
	}
	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	fmt.Println("OUTPUT:")
	fmt.Print("Pekenalkan saya adalah ", nama, ", salah satu mahasiswa prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim, ".")
}
