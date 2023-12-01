package main

import "testing"

func Test_day1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"input", "./input.txt", 53389},
		{"invalid", "", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1(tt.filename); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
