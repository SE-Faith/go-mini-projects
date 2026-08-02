package main

import (
	"fmt"
	"math"
)

func PaymentCalculator(principal float64, rate float64, years float64) float64 {
	annualRate := rate / 100.0
	monthlyRate := annualRate / 12.0

	months := years * 12.0

	monthlyPayment := principal * (monthlyRate * math.Pow(1+monthlyRate, months)) / (math.Pow(1+monthlyRate, months) - 1)

	return monthlyPayment

}

func printSchedule(principal float64, rate float64, years float64, monthlyPayment float64) {

	annualRate := rate / 100.0
	monthlyRate := annualRate / 12.0

	months := years * 12.0
	balance := principal

	fmt.Printf("%-7s %-17s %-17s %-17s \n", "Month", "Interest Paid", "Principal Paid", "Remaining Balance")
	fmt.Println("_____________________________________________________")

	for i := 1.0; i <= months; i++ {
		interestPaid := balance * monthlyRate
		principalPaid := monthlyPayment - interestPaid
		balance -= principalPaid

		if balance < 0 {
			balance = 0
		}
		fmt.Printf("%-10.1f %-14.2f %-14.2f %-14.2f \n ", i, interestPaid, principalPaid, balance)
	}
}

func main() {
	payment := PaymentCalculator(1000.0, 10.0, 1.5)
	fmt.Printf("Monthly Payment is %.2f \n", payment)
	printSchedule(1000, 10, 1.50, 60.06)

}
