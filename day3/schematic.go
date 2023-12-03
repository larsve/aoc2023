package main

import (
	"github.com/larsve/aoc2023/tools"
)

type (
	point struct {
		x, y int
	}
	nbr struct {
		n           int
		topLeft     point
		bottomRight point
	}
	sym struct {
		s  rune
		at point
	}
	schematic struct {
		numbers []nbr
		symbols []sym
	}
)

func parse(input *tools.LineReader) schematic {
	var numbers []nbr
	var symbols []sym
	var y int
	input.ForEach(func(s string) {
		y++
		var n, xStart int
		addNbr := func(xStop int) {
			if n == 0 && xStart == 0 {
				return
			}
			numbers = append(numbers, nbr{n, point{xStart, y - 1}, point{xStop + 1, y + 1}})
			n, xStart = 0, 0
		}
		addSym := func(r rune, x int) {
			symbols = append(symbols, sym{s: r, at: point{x + 1, y}})
		}
		for x, r := range s {
			switch {
			case r >= '0' && r <= '9':
				if n == 0 {
					xStart = x
				}
				n = n*10 + int(r-'0')
			case r == '.':
				addNbr(x)
			default:
				addNbr(x)
				addSym(r, x)
			}
		}
		addNbr(len(s) - 1)
	})
	return schematic{numbers, symbols}
}

func (s sym) adjacent(n nbr) bool {
	if s.at.x < n.topLeft.x || s.at.x > n.bottomRight.x {
		return false
	}
	if s.at.y < n.topLeft.y || s.at.y > n.bottomRight.y {
		return false
	}
	return true
}

func (s schematic) partNumSum() int {
	var pns int
	for _, sym := range s.symbols {
		for _, nbr := range s.numbers {
			if sym.adjacent(nbr) {
				pns += nbr.n
			}
		}
	}
	return pns
}

func (s schematic) gearRatio() int {
	var sum, nbrCnt, nbrSum int
	for _, sym := range s.symbols {
		if sym.s != '*' {
			continue
		}
		nbrCnt, nbrSum = 0, 1
		for _, nbr := range s.numbers {
			if sym.adjacent(nbr) {
				nbrCnt++
				nbrSum *= nbr.n
			}
		}
		if nbrCnt > 1 {
			sum += nbrSum
		}
	}
	return sum
}
