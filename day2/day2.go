package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

func day2(filename string) []int {
	f, close, err := tools.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return nil
	}
	defer close()

	games := parseGames(f)

	return []int{games.getScoreForMaxCubes(12, 13, 14), games.getPower()}
}

func main() {
	fmt.Println("Day2:", day2("./day2/input.txt"))
}
