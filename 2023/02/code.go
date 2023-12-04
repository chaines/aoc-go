package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type CubeSet = map[string]int

var gameSet = CubeSet{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		totalPower := 0
		for _, game := range strings.Split(input, "\n") {
			game = strings.Split(game, ": ")[1]
			totalPower += gamePower(minCubes(game))
		}
		return totalPower
	}

	total := 0
	for i, game := range strings.Split(input, "\n") {
		game = strings.Split(game, ": ")[1]
		if gamePossible(game, gameSet) {
			total += i + 1
		}
	}
	return total
}

func parsePulls(game string) []CubeSet {
	pulls := strings.Split(game, "; ")
	cubeSets := make([]CubeSet, len(pulls))

	for i, pull := range strings.Split(game, "; ") {
		cubeSets[i] = CubeSet{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, colorPull := range strings.Split(pull, ", ") {
			pullVals := strings.Split(colorPull, " ")
			count, color := pullVals[0], pullVals[1]
			countVal, err := strconv.Atoi(count)
			if err != nil {
				panic(err)
			}

			if countVal > cubeSets[i][color] {
				cubeSets[i][color] = countVal
			}
		}
	}
	return cubeSets
}

func gamePossible(game string, gameSet CubeSet) bool {
	for _, cubeSet := range parsePulls(game) {
		if cubeSet["red"] > gameSet["red"] || cubeSet["green"] > gameSet["green"] || cubeSet["blue"] > gameSet["blue"] {
			return false
		}
	}
	return true
}

func gamePower(cubeSet CubeSet) int {
	return cubeSet["red"] * cubeSet["green"] * cubeSet["blue"]
}

func minCubes(game string) CubeSet {
	cubeSet := CubeSet{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	keys := []string{"red", "green", "blue"}
	for _, gameSet := range parsePulls(game) {
		for _, key := range keys {
			if gameSet[key] > cubeSet[key] {
				cubeSet[key] = gameSet[key]
			}
		}
	}
	return cubeSet
}
