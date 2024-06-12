package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const FILE_NAME = "problems.csv"

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println("error opening CSV file")
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading lines")
		os.Exit(1)
	}
	fmt.Println(lines)
}
