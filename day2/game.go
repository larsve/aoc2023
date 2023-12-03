package main

import (
	"fmt"

	"github.com/larsve/aoc2023/tools"
)

type (
	game struct {
		ID            int
		redCubesMax   int
		greenCubesMax int
		blueCubesMax  int
		power         int
	}
	games []game
)

func parseGame(line string) game {
	var g game
	ll := len(line)
	if ll < 13 {
		return g
	}

	var nbr, rc, gc, bc int
	var clr string
	setMax := func(v *int, val int) {
		if *v > val {
			return
		}
		*v = val
	}
	addCubes := func(isSet bool) {
		switch clr {
		case "red":
			rc += nbr
		case "green":
			gc += nbr
		case "blue":
			bc += nbr
		default:
			fmt.Printf("unknown color: %s\n", clr)
		}
		clr = ""
		nbr = 0
		if isSet {
			setMax(&g.redCubesMax, rc)
			setMax(&g.greenCubesMax, gc)
			setMax(&g.blueCubesMax, bc)
			rc, gc, bc = 0, 0, 0
		}
	}
	for i := 5; i < ll; i++ {
		r := line[i]
		switch {
		case r >= '0' && r <= '9':
			nbr = nbr*10 + int(r-'0')
		case r >= 'a' && r <= 'z':
			clr += string(r)
		case r == ':':
			g.ID = nbr
			nbr = 0
		case r == ',' || r == ';':
			addCubes(r == ';')
		default:
		}
	}
	addCubes(true)
	g.power = g.redCubesMax * g.greenCubesMax * g.blueCubesMax
	return g
}

func parseGames(input *tools.LineReader) games {
	var g games
	input.ForEach(func(s string) { g = append(g, parseGame(s)) })
	return g
}

func (gl games) getScoreForMaxCubes(r, g, b int) int {
	var sum int
	for _, game := range gl {
		if game.redCubesMax > r || game.greenCubesMax > g || game.blueCubesMax > b {
			continue
		}
		sum += game.ID
	}
	return sum
}

func (gl games) getPower() int {
	var sum int
	for _, game := range gl {
		sum += game.power
	}
	return sum
}
