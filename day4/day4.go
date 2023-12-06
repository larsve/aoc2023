package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

func day4(filename string) []int {
	f, close, err := tools.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return nil
	}
	defer close()

	cards := parse(f)
	return []int{cards.points(), cards.totalScratchcards()}
}

func main() {
	fmt.Println("Day4:", day4("./day4/input.txt"))
}
