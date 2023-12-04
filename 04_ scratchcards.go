package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func scratchcards(filename string, calcFun func(string, map[int]int) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int
	mapWins := map[int]int{}

	for fileScanner.Scan() {
		text := fileScanner.Text()
		result += calcFun(text, mapWins)
	}

	return result
}

// func scratchcards2(filename string, calcFun func(string, map[int]int) int) int {
// 	f := readFile(filename)
// 	defer closeFile(f)

// 	fileScanner := bufio.NewScanner(f)

// 	var result int
// 	mapWins := map[int]int{}

// 	for fileScanner.Scan() {
// 		text := fileScanner.Text()
// 		calcFun(text, mapWins)
// 	}

// 	return result
// }

func countWorth(text string, mapWins map[int]int) int {
	var res int

	_, winning, my := parseLine(text)
	wins := countWinning(winning, my)
	res = calcScore(wins)

	return res
}

func parseLine(text string) (int, map[int]bool, map[int]bool) {
	split1 := strings.Split(text, ":")
	strCard := strings.Fields(split1[0])[1]
	card, err := strconv.Atoi(strCard)
	if err != nil {
		// log.Print(err)
	}
	split2 := strings.Split(split1[1], "|")
	splitWinning := splitToInt(split2[0])
	splitMy := splitToInt(split2[1])

	fmt.Printf("Card: %d Winning: %v My: %v\n", card, splitWinning, splitMy)

	return card, splitWinning, splitMy
}

func splitToInt(text string) map[int]bool {
	tmp := strings.Split(text, " ")
	values := make(map[int]bool, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			// log.Print(err)
			continue
		}
		values[v] = true
	}
	return values
}

func countWinning(winning map[int]bool, my map[int]bool) int {
	var res int = 0
	// fmt.Printf("Winning: %v My: %v\n", winning, my)

	for k := range my {
		// fmt.Println("Key ", k)
		if winning[k] {
			// fmt.Println("Found ", k)
			res++
		}
	}
	// fmt.Println("Result ", res)
	return res
}

func calcScore(winCount int) int {
	if winCount == 0 {
		return 0
	}

	var res int = 1

	for i := 0; i < winCount-1; i++ {
		res *= 2
	}

	// fmt.Println("Score ", res)
	return res
}

func calcScore2(card int, winCount int, input map[int]int) (map[int]int, int) {
	// fmt.Println("-Card ", card, " cnt: ", input[card])
	input[card] += 1
	// fmt.Println("--Card ", card, " cnt: ", input[card])

	for i := card + 1; i <= card+winCount; i++ {
		input[i] += 1 * input[card]
	}

	// fmt.Println("Input ", input)
	return input, input[card]
}

func totalCards(text string, mapWins map[int]int) int {
	card, winning, my := parseLine(text)
	wins := countWinning(winning, my)

	var res, tmp int = 0, 0

	mapWins, tmp = calcScore2(card, wins, mapWins)

	res += tmp

	return res
}
