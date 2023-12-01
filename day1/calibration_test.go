package main

import (
	"testing"

	"github.com/larsve/aoc2023/tools"
)

var (
	nbrsToiTestCases = []struct {
		name  string
		input string
		want  rune
	}{
		{"1", "one.......", 1},
		{"2", "two.......", 2},
		{"3", "three.....", 3},
		{"4", "four......", 4},
		{"5", "five......", 5},
		{"6", "six.......", 6},
		{"7", "seven.....", 7},
		{"8", "eight.....", 8},
		{"9", "nine......", 9},
		{"*", "..........", 0},
		{"_", "", 0},
	}
	calibrationTestCases = []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example1",
			input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			want:  142,
		},
		{
			name:  "example2",
			input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			want:  281,
		},
		{name: one, input: one, want: 11},
		{name: two, input: two, want: 22},
		{name: three, input: three, want: 33},
		{name: four, input: four, want: 44},
		{name: five, input: five, want: 55},
		{name: six, input: six, want: 66},
		{name: seven, input: seven, want: 77},
		{name: eight, input: eight, want: 88},
		{name: nine, input: nine, want: 99},
	}
)

func runNbrsToiTests(t *testing.T, f func(string) (rune, bool)) {
	for _, tt := range nbrsToiTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := f(tt.input); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_nbrsToi1(t *testing.T) { runNbrsToiTests(t, nbrsToi1) }
func Test_nbrsToi2(t *testing.T) { runNbrsToiTests(t, nbrsToi2) }
func Test_nbrsToi3(t *testing.T) { runNbrsToiTests(t, nbrsToi3) }

func Test_calibrationValue(t *testing.T) {
	for _, tt := range calibrationTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibrationValue(tools.New(tt.input)); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func runNbrsToiBenchAllValues(b *testing.B, f func(string) (rune, bool)) {
	var r rune
	params := [...]string{
		"one.......", "two.......", "three.....", "four......", "five......",
		"six.......", "seven.....", "eight.....", "nine......",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range params {
			r, _ = f(s)
		}
	}
	if r != 9 {
		b.Errorf("got: %d, want: 9", r)
	}
}

func runNbrsToiBenchSingleValue(b *testing.B, str string, want rune, f func(string) (rune, bool)) {
	var r rune
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, _ = f(str)
	}
	if r != want {
		b.Errorf("got: %d, want: %d", r, want)
	}
}

func Benchmark_nbrsToi1(b *testing.B) { runNbrsToiBenchAllValues(b, nbrsToi1) }
func Benchmark_nbrsToi2(b *testing.B) { runNbrsToiBenchAllValues(b, nbrsToi2) }
func Benchmark_nbrsToi3(b *testing.B) { runNbrsToiBenchAllValues(b, nbrsToi3) }

func Benchmark_nbrsToi1_One(b *testing.B) { runNbrsToiBenchSingleValue(b, "one.......", 1, nbrsToi1) }
func Benchmark_nbrsToi2_One(b *testing.B) { runNbrsToiBenchSingleValue(b, "one.......", 1, nbrsToi2) }
func Benchmark_nbrsToi3_One(b *testing.B) { runNbrsToiBenchSingleValue(b, "one.......", 1, nbrsToi3) }

func Benchmark_nbrsToi1_Nine(b *testing.B) { runNbrsToiBenchSingleValue(b, "nine......", 9, nbrsToi1) }
func Benchmark_nbrsToi2_Nine(b *testing.B) { runNbrsToiBenchSingleValue(b, "nine......", 9, nbrsToi2) }
func Benchmark_nbrsToi3_Nine(b *testing.B) { runNbrsToiBenchSingleValue(b, "nine......", 9, nbrsToi3) }

func Benchmark_calibrationValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := tools.New("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
		b.StartTimer()
		if got := calibrationValue(input); got != 281 {
			b.Errorf("got: %v, want: 281", got)
		}
	}
}
