package tools

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: ""},
		{input: "one", want: []string{"one"}},
		{input: "one\ntwo", want: []string{"one", "two"}},
		{input: "one\ntwo\nmany\n", want: []string{"one", "two", "many"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			r := New(tt.input)
			var cnt int
			r.ForEach(func(line string) {
				if line != tt.want[cnt] {
					t.Errorf("ERROR: got line#%d: %s, want: %s", cnt, line, tt.want[cnt])
				}
				cnt++
			})
		})
	}
}

func TestOpen(t *testing.T) {
	tests := []struct {
		input   string
		want    []string
		wantErr bool
	}{
		{input: "./testdata/input0.txt"},
		{input: "./testdata/input1.txt", want: []string{"one"}},
		{input: "./testdata/input2.txt", want: []string{"one", "two"}},
		{input: "./testdata/input3.txt", want: []string{"one", "two", "many"}},
		{input: "", wantErr: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			r, c, err := Open(tt.input)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatalf("ERROR: failed to open file, error: %v (%[1]T)", err)
			}
			defer c()
			var cnt int
			r.ForEach(func(line string) {
				if line != tt.want[cnt] {
					t.Errorf("ERROR: got line#%d: %s, want: %s", cnt, line, tt.want[cnt])
				}
				cnt++
			})
		})
	}
}
