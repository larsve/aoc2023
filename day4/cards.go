package main

import (
	"github.com/larsve/aoc2023/tools"
)

type (
	card struct {
		nbr                int
		noOfWinningNumbers int
		copies             int
	}
	cards []card
)

func parseCard(line string) card {
	var c card
	ll := len(line)
	if ll < 12 {
		return c
	}

	var winningNbr []int
	yourNbrMap := make(map[int]struct{})

	var nbr, nbrType int
	addNbr := func() {
		if nbr == 0 {
			return
		}
		switch nbrType {
		case 0:
			c.nbr = nbr
		case 1:
			winningNbr = append(winningNbr, nbr)
		case 2:
			yourNbrMap[nbr] = struct{}{}
		}
		nbr = 0
	}
	for i := 5; i < ll; i++ {
		r := line[i]
		switch {
		case r >= '0' && r <= '9':
			nbr = nbr*10 + int(r-'0')
		case r == ':':
			addNbr()
			nbrType++
		case r == '|':
			addNbr()
			nbrType++
		case r == ' ':
			addNbr()
		}
	}
	addNbr()
	for _, nbr := range winningNbr {
		if _, ok := yourNbrMap[nbr]; ok {
			c.noOfWinningNumbers++
		}
	}
	return c
}

func parse(input *tools.LineReader) cards {
	var c cards
	input.ForEach(func(s string) { c = append(c, parseCard(s)) })
	return c
}

func (cl cards) points() int {
	var sum int
	for _, c := range cl {
		wc := c.noOfWinningNumbers - 1
		if wc < 0 {
			continue
		}
		sum += 1 << wc
	}
	return sum
}

func (cl cards) totalScratchcards() int {
	var sum int
	scl := len(cl)
	for i, c := range cl {
		sum += 1 + c.copies
		wc := c.noOfWinningNumbers
		if wc == 0 {
			continue
		}
		for j := i + 1; j < scl && j <= i+wc; j++ {
			cl[j].copies += 1 + c.copies
		}
	}
	return sum
}
