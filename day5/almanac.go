package main

import (
	"math"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/larsve/aoc2023/tools"
)

type (
	mapRange struct {
		dstStart int
		srcStart int
		length   int
	}
	aMap    []mapRange
	bySrc   aMap
	almanac struct {
		seeds                 []int
		seedToSoil            aMap
		soilToFertilizer      aMap
		fertilizerToWater     aMap
		waterToLight          aMap
		lightToTemperature    aMap
		temperatureToHumidity aMap
		humidityToLocation    aMap
	}
)

func parse(input *tools.LineReader) *almanac {
	var a almanac
	var cMap *aMap
	input.ForEach(func(s string) {
		switch {
		case s == "":
			return
		case strings.HasPrefix(s, "seeds:"):
			var n int
			for _, c := range s[7:] {
				if c >= '0' && c <= '9' {
					n = n*10 + int(c-'0')
				} else {
					a.seeds = append(a.seeds, n)
					n = 0
				}
			}
			if n > 0 {
				a.seeds = append(a.seeds, n)
			}
		case s == "seed-to-soil map:":
			cMap = &a.seedToSoil
		case s == "soil-to-fertilizer map:":
			cMap = &a.soilToFertilizer
		case s == "fertilizer-to-water map:":
			cMap = &a.fertilizerToWater
		case s == "water-to-light map:":
			cMap = &a.waterToLight
		case s == "light-to-temperature map:":
			cMap = &a.lightToTemperature
		case s == "temperature-to-humidity map:":
			cMap = &a.temperatureToHumidity
		case s == "humidity-to-location map:":
			cMap = &a.humidityToLocation
		default:
			var n, p int
			var r mapRange
			for _, c := range s {
				if c >= '0' && c <= '9' {
					n = n*10 + int(c-'0')
				} else {
					p++
					switch p {
					case 1:
						r.dstStart = n
					case 2:
						r.srcStart = n
					}
					n = 0
				}
			}
			r.length = n
			*cMap = append(*cMap, r)
		}

	})

	sort.Sort(bySrc(a.seedToSoil))
	sort.Sort(bySrc(a.soilToFertilizer))
	sort.Sort(bySrc(a.fertilizerToWater))
	sort.Sort(bySrc(a.waterToLight))
	sort.Sort(bySrc(a.lightToTemperature))
	sort.Sort(bySrc(a.temperatureToHumidity))
	sort.Sort(bySrc(a.humidityToLocation))
	return &a
}

func (m aMap) get(src int) int {
	for _, r := range m {
		if src < r.srcStart {
			return src
		}
		if src >= r.srcStart && src < r.srcStart+r.length {
			return r.dstStart + (src - r.srcStart)
		}
	}
	return src
}

func (m bySrc) Len() int           { return len(m) }
func (m bySrc) Less(i, j int) bool { return m[i].srcStart < m[j].srcStart }
func (m bySrc) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (a almanac) seedLocation(seed int) int {
	v := seed
	for _, m := range []aMap{a.seedToSoil, a.soilToFertilizer, a.fertilizerToWater, a.waterToLight, a.lightToTemperature, a.temperatureToHumidity, a.humidityToLocation} {
		v = m.get(v)
	}
	return v
}

func (a almanac) minLocation() int {
	l := math.MaxInt
	for _, s := range a.seeds {
		v := a.seedLocation(s)
		if v < l {
			l = v
		}
	}
	return l
}

func (a almanac) minLocationOfSeedRanges() int {
	// Pretty slow, lets check how well this scales with more working threads
	l := math.MaxInt
	for i := 0; i < len(a.seeds); i += 2 {
		start := a.seeds[i]
		length := a.seeds[i+1]
		for j := 0; j < length; j++ {
			v := a.seedLocation(start + j)
			if v < l {
				l = v
			}
		}
	}
	return l
}

func (a almanac) parallellMinLocationOfSeedRanges() int {
	// Faster, completes in around 30 seconds now..
	l := math.MaxInt
	var workWg, resWg sync.WaitGroup
	workChan := make(chan []int, 1)
	resChan := make(chan int, 1)
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		workWg.Add(1)
		go func() {
			defer workWg.Done()
			for sr := range workChan {
				l := math.MaxInt
				for _, s := range sr {
					v := a.seedLocation(s)
					if v < l {
						l = v
					}
				}
				resChan <- l
			}
		}()
	}
	resWg.Add(1)
	go func() {
		defer resWg.Done()
		for v := range resChan {
			if v < l {
				l = v
			}
		}
	}()
	var wb []int
	for i := 0; i < len(a.seeds); i += 2 {
		start := a.seeds[i]
		length := a.seeds[i+1]
		for j := 0; j < length; j++ {
			wb = append(wb, start+j)
			if len(wb) == 1000 {
				workChan <- wb
				wb = nil
			}
		}
	}
	if len(wb) > 0 {
		workChan <- wb
	}
	close(workChan)
	workWg.Wait()
	close(resChan)
	resWg.Wait()
	return l
}
