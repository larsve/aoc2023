package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

func day5(filename string) []int {
	f, close, err := tools.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return nil
	}
	defer close()

	almanac := parse(f)
	return []int{almanac.minLocation(), almanac.parallellMinLocationOfSeedRanges()}
}

func main() {
	fmt.Println("Day5:", day5("./day5/input.txt"))
}
