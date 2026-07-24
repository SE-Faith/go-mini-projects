package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*

PROGRAM FLOW

1. Display Options
2. Read choice
3. Prompt for value
4. Clean and validate input
5. Call function
6. Display output
*/

func main() {

	fmt.Print("Hello! \n")
	fmt.Print("Pick the conversion type \n")
	fmt.Print("1. CelsiusToFahrenheit \n")
	fmt.Print("2. FahrenheitToCelsius \n")
	fmt.Print("3. CelsiusToKelvin \n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	converted, err := strconv.Atoi(input)

	if err != nil {
		fmt.Print(err, "\n Please enter a valid number")
		return
	}

	fmt.Print("Enter the temperature to convert \n")
	reader = bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number)
	numbered, err := strconv.ParseFloat(number, 64)

	if err != nil {
		fmt.Print(err, "\n Please enter a valid number(")

	}

	if converted == 1 {

		f := CelsiusToFahrenheit(numbered)
	fmt.Println("temperature: ", f)
		
	} else if converted == 2 {
		c := FahrenheitToCelsius(numbered)
		fmt.Println("temperature: ", c)
	} else if converted == 3 {
		k := CelsiusToKelvin(numbered)
		fmt.Println("temperature: ", k)
	} else {
		fmt.Println("You have entered an invalid value")
	}

}

func CelsiusToFahrenheit(c float64) float64{

	f := (c * 9.0 / 5.0) + 32.0
	return f

}

func FahrenheitToCelsius(f float64) float64 {
	c := (f - 32.0) * 5.0 / 9.0
	return c
}

func CelsiusToKelvin(c float64) float64 {
	k := c + 273.15
	return k
}
