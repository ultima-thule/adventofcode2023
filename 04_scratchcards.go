package main

import (
	"strconv"
	"strings"
)

// Solve puzzle no 1
func countWorth(text string, mapWins map[int]int) int {
	_, myWinning := parseLine(text)

	return calcScorePart1(len(myWinning))
}

// Parse single input line and return list of winning numbers
func parseLine(text string) (int, map[int]bool) {
	split1 := strings.Split(text, ":")
	card, _ := strconv.Atoi(strings.Fields(split1[0])[1])

	split2 := strings.Split(split1[1], "|")
	winning := splitToInt(split2[0], nil)
	myWinning := splitToInt(split2[1], winning)

	return card, myWinning
}

// Split text into int values and add to result only those which are present in reference map
func splitToInt(text string, winning map[int]bool) map[int]bool {
	tmp := strings.Fields(text)
	values := make(map[int]bool, len(tmp))

	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			// log.Print(err)
			continue
		}
		if winning == nil || winning[v] {
			values[v] = true
		}
	}
	return values
}

// Calculate score of winning cards
func calcScorePart1(winCount int) int {
	if winCount == 0 {
		return 0
	}

	var res int = 1
	for i := 0; i < winCount-1; i++ {
		res *= 2
	}
	return res
}

// Add cards based on no of wins
func calcScorePart2(card int, winCount int, input map[int]int) (map[int]int, int) {
	input[card] += 1

	for i := card + 1; i <= card+winCount; i++ {
		input[i] += 1 * input[card]
	}
	return input, input[card]
}

// Solve puzzle no 2
func countTotalCards(text string, mapWins map[int]int) int {
	card, myWinning := parseLine(text)

	var res, tmp int = 0, 0
	mapWins, tmp = calcScorePart2(card, len(myWinning), mapWins)
	res += tmp

	return res
}
