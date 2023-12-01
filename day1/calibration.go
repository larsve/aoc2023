package main

import (
	"strings"

	"github.com/larsve/aoc2023/tools"
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var nbrs = [...]string{one, two, three, four, five, six, seven, eight, nine}

// nbrsToi1 is small by relatively slow
func nbrsToi1(s string) (rune, bool) {
	for i, n := range nbrs {
		if strings.HasPrefix(s, n) {
			return rune(i + 1), true
		}
	}
	return 0, false
}

// nbrsToi2 is pretty fast, pretty low cognitive complexity
func nbrsToi2(s string) (rune, bool) {
	if len(s) < 3 {
		return 0, false
	}
	c := s[0]
	switch {
	case c == 'o' && strings.HasPrefix(s, one):
		return 1, true
	case c == 't' && strings.HasPrefix(s, two):
		return 2, true
	case c == 't' && strings.HasPrefix(s, three):
		return 3, true
	case c == 'f' && strings.HasPrefix(s, four):
		return 4, true
	case c == 'f' && strings.HasPrefix(s, five):
		return 5, true
	case c == 's' && strings.HasPrefix(s, six):
		return 6, true
	case c == 's' && strings.HasPrefix(s, seven):
		return 7, true
	case c == 'e' && strings.HasPrefix(s, eight):
		return 8, true
	case c == 'n' && strings.HasPrefix(s, nine):
		return 9, true
	default:
		return 0, false
	}
}

// nbrsToi3 is the fastest of these three, but have a higher cognitive complexity and many lines of code
func nbrsToi3(s string) (rune, bool) {
	if len(s) < 3 {
		return 0, false
	}
	switch s[0] {
	case 'o':
		if strings.HasPrefix(s, one) {
			return 1, true
		}
	case 't':
		if strings.HasPrefix(s, two) {
			return 2, true
		}

		if strings.HasPrefix(s, three) {
			return 3, true
		}
	case 'f':
		if strings.HasPrefix(s, four) {
			return 4, true
		}
		if strings.HasPrefix(s, five) {
			return 5, true
		}
	case 's':
		if strings.HasPrefix(s, six) {
			return 6, true
		}
		if strings.HasPrefix(s, seven) {
			return 7, true
		}
	case 'e':
		if strings.HasPrefix(s, eight) {
			return 8, true
		}
	case 'n':
		if strings.HasPrefix(s, nine) {
			return 9, true
		}
	default:
		return 0, false
	}
	return 0, false
}

func calibrationValue(input *tools.LineReader) int64 {
	var cv int64
	var first, last rune
	var firstOk, ok bool
	input.ForEach(func(s string) {
		firstOk = false
		for ofs, c := range s {
			if c >= '0' && c <= '9' {
				c = c - '0'
			} else {
				c, ok = nbrsToi3(s[ofs:])
				if !ok {
					continue
				}
			}
			if !firstOk {
				first = c
				firstOk = true
			}
			last = c
		}
		cv += int64(first*10 + last)
	})
	return cv
}
