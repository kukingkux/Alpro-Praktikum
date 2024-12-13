package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Enter a word: ")
	_, err := fmt.Scan(&input) // Read a single word
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Print each character in the input string
	for _, r := range input {
		fmt.Printf("%c ", r)
	}
	fmt.Println() // Print a newline for better formatting
}
