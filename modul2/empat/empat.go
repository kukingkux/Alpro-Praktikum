package main

import "fmt"

func main() {
	var (
		fahrenheit int
		celcius    float64
	)
	fmt.Print("Masukkan suhu dalam fahrenheit: ")
	fmt.Scan(&fahrenheit)
	var floatFahrenheit float64 = float64(fahrenheit)
	celcius = (floatFahrenheit - 32) * 5 / 9
	fmt.Println("Suhu dalam celcius adalah: ", celcius)
}
