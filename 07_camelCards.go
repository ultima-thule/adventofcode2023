package main

import (
	"fmt"
	"sort"
	"strings"
)

// Day 04 solution
func camelCards(filename string, calcFun func([]game) int, prepFun func(data []string) []game) int {
	input := readInput07(filename, calcFun, prepFun)
	return calcFun(input)
}

// Solve puzzle no 1
func draw(input []game) int {
	if input == nil || len(input) < 2 {
		return 0
	}

	var res int = 0

	for i := 0; i < len(input); i++ {
		input[i].power = calcDrawPower(input[i].hand, input[i].bid)
	}

	sortCards(input)

	for i := 0; i < len(input); i++ {
		res += input[i].bid * (i + 1)
	}

	return res
}

// Solve puzzle no 2
func draw2(input []game) int {
	if input == nil || len(input) < 2 {
		return 0
	}

	var res int = 0

	for i := 0; i < len(input); i++ {
		input[i].power = calcDrawPower2(input[i].hand, input[i].bid)
	}

	sortCards(input)

	for i := 0; i < len(input); i++ {
		res += input[i].bid * (i + 1)
	}

	return res
}

func calcDrawPower(cards string, bid int) int {
	hand := map[rune]int{}

	for _, v := range cards {
		hand[v] += 1
	}

	// fmt.Println("Hand: ", cards, " bid ", bid, " cnt ", len(hand))

	switch len(hand) {
	case 1:
		return 7
	case 2:
		if checkForValue(4, hand) {
			return 6
		}
		return 5
	case 3:
		if checkForValue(3, hand) {
			return 4
		}
		return 3
	case 4:
		return 2
	case 5:
		return 1
	}
	return 0
}

func calcDrawPower2(cards string, bid int) int {
	hand := map[rune]int{}

	for _, v := range cards {
		hand[v] += 1
	}

	fmt.Println("\nHand: ", cards, " bid ", bid, " cnt ", len(hand))
	var jokerCnt int = hand['J']
	fmt.Println("Joker count: ", jokerCnt)
	delete(hand, 'J')
	fmt.Println("Hand: ", cards, " bid ", bid, " cnt", len(hand))

	if jokerCnt > 0 {
		var maxCard int
		var maxKey rune
		for k, v := range hand {
			if v > maxCard {
				maxCard = v
				maxKey = k
			}
		}

		fmt.Println("Found max card: ", maxCard, " at ", string(maxKey), "\n")
		hand[maxKey] += jokerCnt
		fmt.Println(hand, "\n")
	}

	switch len(hand) {
	case 1:
		return 7
	case 2:
		if checkForValue(4, hand) {
			return 6
		}
		return 5
	case 3:
		if checkForValue(3, hand) {
			return 4
		}
		return 3
	case 4:
		return 2
	case 5:
		return 1
	}
	return 0
}

func checkForValue(value int, deck map[rune]int) bool {
	for _, v := range deck {
		if v == value {
			return true
		}
	}
	return false
}

type game struct {
	hand  string
	bid   int
	power int
}

func prepData07(input []string) []game {
	if input == nil || len(input) < 2 {
		return nil
	}

	var tmp []string
	output := make([]game, 0)

	for i := 0; i < len(input); i++ {
		tmp = strings.Split(input[i], " ")
		output = append(output, game{tmp[0], convert(tmp[1]), 0})
	}

	return output
}

func sortCards(sl []game) {
	sort.Slice(sl, func(i, j int) bool {
		if sl[i].power != sl[j].power {
			// fmt.Println("Different powers")
			return sl[i].power < sl[j].power
		}
		var res bool = false
		// fmt.Println("Compare: ", sl[i], " ", sl[j])
		// fmt.Println("Same powers")
		for k := 0; k < len(sl[i].hand); k++ {
			// fmt.Println("Analysing letter ", sl[i].hand[k], " and ", sl[j].hand[k])
			if sl[i].hand[k] != sl[j].hand[k] {
				return compareCards(sl[i].hand[k], sl[j].hand[k])
			}
		}
		return res
	})
}

func sortCards2(sl []game) {
	sort.Slice(sl, func(i, j int) bool {
		if sl[i].power != sl[j].power {
			// fmt.Println("Different powers")
			return sl[i].power < sl[j].power
		}
		var res bool = false
		// fmt.Println("Compare: ", sl[i], " ", sl[j])
		// fmt.Println("Same powers")
		for k := 0; k < len(sl[i].hand); k++ {
			// fmt.Println("Analysing letter ", sl[i].hand[k], " and ", sl[j].hand[k])
			if sl[i].hand[k] != sl[j].hand[k] {
				return compareCards2(sl[i].hand[k], sl[j].hand[k])
			}
		}
		return res
	})
}

func compareCards(r1 byte, r2 byte) bool {
	mapPower := map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}
	// fmt.Println("Comparing letters ", string(r1), " and ", string(r2), " result ", mapPower[string(r1)], mapPower[string(r2)])
	return mapPower[string(r1)] > mapPower[string(r2)]
}

func compareCards2(r1 byte, r2 byte) bool {
	mapPower := map[string]int{"A": 1, "K": 2, "Q": 3, "T": 4, "9": 5, "8": 6, "7": 7, "6": 8, "5": 9, "4": 10, "3": 11, "2": 12, "J": 13}
	// fmt.Println("Comparing letters ", string(r1), " and ", string(r2), " result ", mapPower[string(r1)], mapPower[string(r2)])
	return mapPower[string(r1)] > mapPower[string(r2)]
}
