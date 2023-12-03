package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/larsve/aoc2023/tools"
)

const exampleInput = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"

func Test_parseGame(t *testing.T) {
	tests := []struct {
		input string
		want  game
	}{
		{"Game 1: 3 rd", game{0, 0, 0, 0, 0}},
		{"Game 1: 1 red", game{1, 1, 0, 0, 0}},
		{"Game 1: 2 green", game{1, 0, 2, 0, 0}},
		{"Game 1: 3 blue", game{1, 0, 0, 3, 0}},
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", game{1, 4, 2, 6, 4 * 2 * 6}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", game{2, 1, 3, 4, 1 * 3 * 4}},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", game{3, 20, 13, 6, 20 * 13 * 6}},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", game{4, 14, 3, 15, 14 * 3 * 15}},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", game{5, 6, 3, 2, 6 * 3 * 2}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseGame(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_parseGames(t *testing.T) {
	tests := []struct {
		input string
		want  []game
	}{
		{
			input: exampleInput,
			want:  games{{1, 4, 2, 6, 4 * 2 * 6}, {2, 1, 3, 4, 1 * 3 * 4}, {3, 20, 13, 6, 20 * 13 * 6}, {4, 14, 3, 15, 14 * 3 * 15}, {5, 6, 3, 2, 6 * 3 * 2}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			got := parseGames(tools.New(tt.input))
			if len(got) != len(tt.want) {
				t.Fatalf("got %d games, want %d games", len(got), len(tt.want))
			}
			for i, g := range got {
				if !reflect.DeepEqual(g, tt.want[i]) {
					t.Errorf("got: %v, want: %v", g, tt.want[i])
				}
			}
		})
	}
}

func Test_getScoreForMaxCubes(t *testing.T) {
	tests := []struct {
		input *tools.LineReader
		r     int
		g     int
		b     int
		want  int
	}{
		{tools.New(exampleInput), 12, 13, 14, 8},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			g := parseGames(tt.input)
			if got := g.getScoreForMaxCubes(tt.r, tt.g, tt.b); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_games_getPower(t *testing.T) {
	tests := []struct {
		input *tools.LineReader
		want  int
	}{
		{tools.New(exampleInput), 2286},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_#%d", i), func(t *testing.T) {
			g := parseGames(tt.input)
			if got := g.getPower(); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_parseGame(b *testing.B) {
	var got game
	for i := 0; i < b.N; i++ {
		got = parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	}
	b.StopTimer()
	want := game{1, 4, 2, 6, 4 * 2 * 6}
	if !reflect.DeepEqual(got, want) {
		b.Errorf("got: %v, want: %v", got, want)
	}
}

func Benchmark_parseGames(b *testing.B) {
	var got games
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := tools.New(exampleInput)
		b.StartTimer()
		got = parseGames(input)
	}
	if got == nil {
		b.Errorf("parseGames failed")
	}
}

func Benchmark_getScoreForMaxCubes(b *testing.B) {
	games := parseGames(tools.New(exampleInput))
	var got int
	want := 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got = games.getScoreForMaxCubes(12, 13, 14)
	}
	if got != want {
		b.Errorf("got: %d, want: %d", got, want)
	}
}

func Benchmark_getPower(b *testing.B) {
	games := parseGames(tools.New(exampleInput))
	var got int
	want := 2286
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got = games.getPower()
	}
	if got != want {
		b.Errorf("got: %d, want: %d", got, want)
	}
}
