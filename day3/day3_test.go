package main

import (
	"reflect"
	"testing"
)

func Test_day3(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []int
	}{
		{"input", "./input.txt", []int{531932, 73646890}},
		{"invalid", "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3(tt.filename)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
