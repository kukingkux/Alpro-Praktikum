package main

import "fmt"

func main() {
	var namatanaman string
	fmt.Scan(&namatanaman)

	switch namatanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Tidak Asli Indonesia")
	default:
		fmt.Println("Tidak termauk tanaman karnivora")
	}
}
