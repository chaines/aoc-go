package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Number struct {
  Seen bool
  Value int
  Start int
  End int
  Row int
}

type Point struct {
  Row int
  Col int
  Val rune
}

const GEAR rune = '*'

var ADJACENT_TILES = [][]int{
  {0, -1},
  {0, 1},
  {-1, -1},
  {-1, 0},
  {-1, 1},
  {1, -1},
  {1, 0},
  {1, 1},
}

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
  numbers, specialChars := parseInput(input)
  if part2 {
    gearRatios := findGearRatios(numbers, specialChars)
    total := 0
    for _, num := range gearRatios {
      total += num
    }
    return total
  }
  partNumbers := findPartNumbers(numbers, specialChars)
  total := 0
  for _, num := range partNumbers {
    total += num.Value
  }
  return total
}

func parseInput(input string) ([]Number, []Point) {
  numbers := make([]Number, 0, 10000)
  specialChars := make([]Point, 0, 10000)

  for row, line := range strings.Split(input, "\n") {
    col := 0
    for col < len(line) {
      if (line[col] >= '0' && line[col] <= '9') {
        // We got us a number
        start := col
        val := ""
        for col < len(line) && line[col] >= '0' && line[col] <= '9' {
          val = val + string(line[col])
          col++
        }

        num, err := strconv.Atoi(val)
        if err != nil {
          panic(err)
        }
        numbers = append(numbers, Number{
          Value: num,
          Start: start,
          End: col - 1,
          Row: row,
          Seen: false,
        })
        continue
      } else if line[col] != '.' {
        specialChars = append(specialChars, Point{
          Row: row,
          Col: col,
          Val: rune(line[col]),
        })
        col++
      } else {
        col++
      }
    }
  }
  return numbers, specialChars
}

func numContainsPoint(number Number, row int, col int) bool {
  return number.Row == row && number.Start <= col && number.End >= col;
}

func findPartNumbers(numbers []Number, specialChars []Point) []Number {
  partNumbers := make([]Number, 0)

  for _, char := range specialChars {
    for _, adj := range ADJACENT_TILES {
      row, col := char.Row + adj[0], char.Col + adj[1]
      for _, num := range numbers {
        if slices.Contains(partNumbers, num) {
          continue
        }
        if numContainsPoint(num, row, col) {
          partNumbers = append(partNumbers, num)
        }
      }
    }
  }
  return partNumbers
}

func findGearRatios(numbers []Number, specialChars []Point) []int {
  gearRatios := make([]int, 0, len(specialChars))
  numbersSeen := make([]Number, 0, len(numbers))

  for _, char := range specialChars {
    if char.Val != GEAR {
      continue
    }

    count := 0
    adjacentNumbers := make([]Number, 6)

    for _, adj := range ADJACENT_TILES {
      row, col := char.Row + adj[0], char.Col + adj[1]
      for _, num := range numbers {
        if slices.Contains(numbersSeen, num) {
          continue;
        }
        if numContainsPoint(num, row, col) {
          adjacentNumbers[count] = num;
          numbersSeen = append(numbersSeen, num)
          count++;
        }
      }
    }
    if count == 2 {
      gearRatios = append(gearRatios, adjacentNumbers[0].Value * adjacentNumbers[1].Value)
    }
  }

  return gearRatios
}

