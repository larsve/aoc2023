package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/larsve/aoc2023/tools"
)

const exampleInput = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"

func Test_parseCard(t *testing.T) {
	tests := []struct {
		line string
		want card
	}{
		{"", card{}},
		{"Card 1: 2 | 3", card{1, 0, 0}},
		{"Card 2:  3 |  4", card{2, 0, 0}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", card{3, 2, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			if got := parseCard(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_cards_points(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: exampleInput, want: 13},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			cl := parse(tools.New(tt.input))
			if got := cl.points(); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_cards_totalScratchcards(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: exampleInput, want: 30},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			cl := parse(tools.New(tt.input))
			if got := cl.totalScratchcards(); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
