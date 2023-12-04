package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
  sum := 0
  for _, line := range strings.Split(input, "\n") {
    if line == "" {
      fmt.Println("Input: ");
      fmt.Println(input);
    }
    sum += Join(FirstNumber(line, part2), LastNumber(line, part2));
  }
  return sum;
}

func Join(first string, second string) int {
  num, err := strconv.Atoi(first + second);
  if err != nil {
    panic(err);
  }
  return num;
}

var numWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};

func FirstNumber(line string, checkWords bool) string {
  for i := range line {
    char := line[i:i+1];
    _, err := strconv.Atoi(char);
    if err == nil {
      return char;
    }

    if !checkWords {
      continue;
    }

    substr := line[0:i+1];
    for i, n := range numWords {
      if strings.Contains(substr, n) {
        return strconv.Itoa(i + 1);
      }
    }
  }
  return "";
}

func LastNumber(line string, checkWords bool) string {
  for i := range line {
    j := len(line) - i - 1;
    char := line[j:j+1];
    _, err := strconv.Atoi(char);
    if err == nil {
      return char;
    }

    if !checkWords {
      continue;
    }

    substr := line[j:];
    for i, n := range numWords {
      if strings.Contains(substr, n) {
        return strconv.Itoa(i + 1);
      }
    }
  }
  return "";
}

