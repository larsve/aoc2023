package main

import (
	"reflect"
	"testing"

	"github.com/larsve/aoc2023/tools"
)

const exampleInput = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

func Test_parse(t *testing.T) {
	tests := []struct {
		name  string
		input *tools.LineReader
		want  schematic
	}{
		{"example", tools.New(exampleInput), schematic{
			numbers: []nbr{
				{467, point{0, 0}, point{4, 2}},
				{114, point{5, 0}, point{9, 2}},
				{35, point{2, 2}, point{5, 4}},
				{633, point{6, 2}, point{10, 4}},
				{617, point{0, 4}, point{4, 6}},
				{58, point{7, 5}, point{10, 7}},
				{592, point{2, 6}, point{6, 8}},
				{755, point{6, 7}, point{10, 9}},
				{664, point{1, 9}, point{5, 11}},
				{598, point{5, 9}, point{9, 11}},
			},
			symbols: []sym{
				{'*', point{4, 2}},
				{'#', point{7, 4}},
				{'*', point{4, 5}},
				{'+', point{6, 6}},
				{'$', point{4, 9}},
				{'*', point{6, 9}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_schematic_partNumSum(t *testing.T) {
	tests := []struct {
		name  string
		input *tools.LineReader
		want  int
	}{
		{"example", tools.New(exampleInput), 4361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parse(tt.input)
			if got := s.partNumSum(); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_schematic_gearRatio(t *testing.T) {
	tests := []struct {
		name  string
		input *tools.LineReader
		want  int
	}{
		{"example", tools.New(exampleInput), 467835},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parse(tt.input)
			if got := s.gearRatio(); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
