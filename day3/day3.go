package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

func day3(filename string) []int {
	f, close, err := tools.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return nil
	}
	defer close()

	schema := parse(f)
	return []int{schema.partNumSum(), schema.gearRatio()}
}

func main() {
	fmt.Println("Day3:", day3("./day3/input.txt"))
}
