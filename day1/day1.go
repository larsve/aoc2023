package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

func day1(filename string) int {
	f, close, err := tools.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return -1
	}
	defer close()

	return int(calibrationValue(f))
}

func main() {
	fmt.Println("Day1:", day1("./day1/input.txt"))
}
