package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <fileName> <logLevel>")
		return
	}
	filePath := os.Args[1]
	logLevel := strings.ToUpper(os.Args[2])
	logLevel = "[" + logLevel + "]"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	var count int = 0
	for reader.Scan() {
		line := reader.Text()

		if strings.Contains(line, logLevel) {
			count++
			fmt.Println(line)

		}

	}
	fmt.Println("Total Count: ", count)
	if reader.Err() != nil {
		log.Fatal(reader.Err())
	}
}
