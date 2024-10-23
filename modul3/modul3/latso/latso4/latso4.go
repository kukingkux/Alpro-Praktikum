package main

import "fmt"

func main() {
	var (
		celcius, reamur, fahrenheit, kelvin float64
	)
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&celcius)
	// celcius = (fahrenheit - 32) * (5.0 / 9.0)
	fahrenheit = (celcius * (9.0 / 5.0)) + 32
	reamur = celcius * (4.0 / 5.0)
	kelvin = (fahrenheit + 459.67) * 5 / 9
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
