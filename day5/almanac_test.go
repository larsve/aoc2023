package main

import (
	"sort"
	"testing"

	"github.com/larsve/aoc2023/tools"
)

const (
	example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
)

func Test_aMap_get(t *testing.T) {
	seedToSoilExample := aMap{mapRange{50, 98, 2}, mapRange{52, 50, 48}}
	sort.Sort(bySrc(seedToSoilExample))
	tests := []struct {
		name string
		m    aMap
		v    int
		want int
	}{
		{name: "low, unmapped", m: seedToSoilExample, v: 49, want: 49},
		{name: "low, mapped 0", m: seedToSoilExample, v: 50, want: 52},
		{name: "low, mapped 1", m: seedToSoilExample, v: 97, want: 99},
		{name: "high, mapped 0", m: seedToSoilExample, v: 98, want: 50},
		{name: "high, mapped 1", m: seedToSoilExample, v: 99, want: 51},
		{name: "high, unmapped", m: seedToSoilExample, v: 100, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.get(tt.v); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_almanac_seedLocation(t *testing.T) {
	a := parse(tools.New(example))
	tests := []struct {
		name string
		a    almanac
		seed int
		want int
	}{
		{name: "seed_79", a: *a, seed: 79, want: 82},
		{name: "seed_14", a: *a, seed: 14, want: 43},
		{name: "seed_55", a: *a, seed: 55, want: 86},
		{name: "seed_13", a: *a, seed: 13, want: 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.seedLocation(tt.seed); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_almanac_minLocation(t *testing.T) {
	a := parse(tools.New(example))
	got := a.minLocation()
	want := 35
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func Test_almanac_minLocationOfSeedRanges(t *testing.T) {
	a := parse(tools.New(example))
	got := a.minLocationOfSeedRanges()
	want := 46
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
