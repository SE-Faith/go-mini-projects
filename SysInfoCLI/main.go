package main

import (
	"fmt"
	"log"
	"os"
	"time"

)

func main() {

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	
	environ := os.Getenv("USERNAME")

	system := os.Getenv("OS")
	
	
	currentTime := time.Now()
	formatted := currentTime.Format("2006-01-02 15:04:05")

	

	var memoryInByte int64 = 17171771717717
	memoryInGB := float64(memoryInByte)/ (1024 * 1024 * 1024)
	

	fmt.Printf("Hostname: %s | Username: %s | OS: %s | Time: %s | Memory: %.2f GB", host, environ, system, formatted, memoryInGB)

}
