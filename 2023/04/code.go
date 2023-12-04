package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
    return runPart2(input)
	}
	return calcScores(input)
}

func runPart2(input string) int {
	cards := strings.Split(input, "\n")
	cardCounts := make([]int, len(cards), len(cards))
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for num, card := range cards {
		// Drop the "Card X: " leading text
		card = strings.Split(card, ": ")[1]
		numStrings := strings.Split(card, " | ")
		winningStrings, cardStrings := numStrings[0], numStrings[1]
		winningNums := trimStrings(chunkString(winningStrings, 3), " ")
		cardNums := trimStrings(chunkString(cardStrings, 3), " ")

		cardWins := 0
		for _, winningNum := range winningNums {
			for _, cardNum := range cardNums {
				if winningNum == cardNum {
					cardWins++
				}
			}
		}
    for i := 1; i <= cardWins; i++ {
      cardCounts[num + i] += cardCounts[num];
    }
	}

	total := 0
	for _, count := range cardCounts {
		total += count
	}
	return total
}

func calcScores(input string) int {
	cards := strings.Split(input, "\n")
	total := 0
	for _, card := range cards {
		// Drop the "Card X: " leading text
		card = strings.Split(card, ": ")[1]
		numStrings := strings.Split(card, " | ")
		winningStrings, cardStrings := numStrings[0], numStrings[1]
		winningNums := trimStrings(chunkString(winningStrings, 3), " ")
		cardNums := trimStrings(chunkString(cardStrings, 3), " ")

		cardScore := 0
		for _, winningNum := range winningNums {
			for _, cardNum := range cardNums {
				if winningNum == cardNum {
					if cardScore == 0 {
						cardScore = 1
					} else {
						cardScore = cardScore * 2
					}
				}
			}
		}
		total += cardScore
	}
	return total
}

func chunkString(input string, chunkSize int) []string {
	chunks := make([]string, 0, (len(input)-1)/chunkSize+1)
	currLen := 0
	currStart := 0
	for i := range input {
		if currLen == chunkSize {
			chunks = append(chunks, input[currStart:i])
			currStart = i
			currLen = 0
		}
		currLen++
	}
	chunks = append(chunks, input[currStart:])
	return chunks
}

func trimStrings(input []string, trimChar string) []string {
	for i := range input {
		input[i] = strings.Trim(input[i], trimChar)
	}
	return input
}
