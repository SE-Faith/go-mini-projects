package main

import (
	"StringSanitizer/sanitizer"
	"fmt"
)

func main() {
	var word string
	fmt.Print("Enter a string to clean: ")
	fmt.Scanln(&word)
	cleanWord := sanitizer.CleanString(word)
	fmt.Printf("The string is %s", cleanWord)
}
