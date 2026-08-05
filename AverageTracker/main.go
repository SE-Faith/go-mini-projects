package main

import (
	"fmt"
	"go-mini-projects/AverageTracker/tracker"
)

func main() {
	readings := []float64{10.0, 12.2, 24.0, 33.9}

	for _, reading := range readings {
		tracker.AddValue(reading)
		fmt.Println("Average is: ", tracker.GetAverage())
	}
}
