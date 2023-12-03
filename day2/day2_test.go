package main

import (
	"reflect"
	"testing"
)

func Test_day2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []int
	}{
		{"input", "./input.txt", []int{2541, 66016}},
		{"invalid", "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day2(tt.filename)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
