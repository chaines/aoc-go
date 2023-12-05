package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Maps struct {
	SeedToSoil         map[int]int
	SoilToFertilizer   map[int]int
	FertilizerToWater  map[int]int
	WaterToLight       map[int]int
	LightToTemp        map[int]int
	TempToHumidity     map[int]int
	HumidityToLocation map[int]int
}

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	seeds, maps := parseInput(input)
	sourceRanges := make([][3]int, 0, len(seeds))
	if part2 {
		for i := 0; i < len(seeds); i += 2 {
			sourceRanges = append(sourceRanges, [3]int{
        seeds[i],
        seeds[i] + seeds[i+1] - 1,
        0,
      })
		}
	} else {

		for _, seed := range seeds {
			sourceRanges = append(sourceRanges, [3]int{seed, seed, 0})
		}
	}

	dests := make([]int, 0, len(sourceRanges))

	for len(sourceRanges) != 0 {
		sourceRange := sourceRanges[0]
		sourceRanges = sourceRanges[1:]

		minSource, maxSource, startMap := sourceRange[0], sourceRange[1], sourceRange[2]
    
    for i := startMap; i < len(maps); i++ {
      mapList := maps[i]
			rangeSolved := false
			for _, mapRange := range mapList {
				minMapSource, maxMapSource := mapRange[2], mapRange[3]
				minMapDest := mapRange[0]

				if minSource < minMapSource && maxSource >= minMapSource {
					// Put it on the queue to deal with later
					sourceRanges = append(sourceRanges, [3]int{minMapSource, maxSource, i})
					maxSource = minMapSource - 1
				}
				if minSource >= minMapSource && minSource <= maxMapSource {
					if maxSource > maxMapSource {
						// Put it on the queue to deal with later
						sourceRanges = append(sourceRanges, [3]int{maxMapSource + 1, maxSource, i})
						maxSource = maxMapSource
					}

					length := maxSource - minSource
					diff := minSource - minMapSource
					minSource = minMapDest + diff
					maxSource = minSource + length
					rangeSolved = true
					break
				}
			}
			if rangeSolved {
				continue
			}
		}
		dests = append(dests, minSource)
	}
	minLocation := math.MaxInt
	for _, dest := range dests {
		if dest < minLocation {
			minLocation = dest
		}
	}
  fmt.Println(minLocation)
	return minLocation
}

func parseInput(input string) ([]int, [][][4]int) {
	// Assume they go in order, because I'm not interested in parsing names
	// as well.

	chunks := strings.Split(input, "\n\n")
	seedNums := strings.Split(chunks[0], ": ")
	seedNums = strings.Split(seedNums[1], " ")
	seeds := make([]int, len(seedNums))
	for i, seed := range seedNums {
		seeds[i] = atoi(seed)
	}

	chunks = chunks[1:]
	maps := make([][][4]int, len(chunks))

	for i, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		lines = lines[1:] // Discard title line, who needs it.
		for _, line := range lines {
			numStrings := strings.Split(line, " ")
			minDest, minSource, length := atoi(numStrings[0]), atoi(numStrings[1]), atoi(numStrings[2])
			nums := [4]int{
				minDest,
				minDest + length - 1,
				minSource,
				minSource + length - 1,
			}
			maps[i] = append(maps[i], nums)
		}
	}

	return seeds, maps
}

func atoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return val
}
