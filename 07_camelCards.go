package main

import (
	"sort"
)

type game struct {
	hand  string
	bid   int
	power int
}

// Solve puzzle no 1 && 2
func puzzle07(input []game, useJoker bool) int {
	if input == nil {
		return 0
	}

	var res int = 0

	for i := 0; i < len(input); i++ {
		input[i].power = calcDrawPower(input[i].hand, input[i].bid, useJoker)
	}

	sortCards(input, useJoker)

	for i := 0; i < len(input); i++ {
		res += input[i].bid * (i + 1)
	}

	return res
}

// calculate power of deck
func calcDrawPower(cards string, bid int, useJoker bool) int {
	hand := map[rune]int{}

	for _, v := range cards {
		hand[v] += 1
	}

	if useJoker {
		var jokerCnt int = hand['J']

		if jokerCnt > 0 && jokerCnt < 55 {
			delete(hand, 'J')
			hand[findOptimumCard(hand)] += jokerCnt
		}
	}

	// map priorities
	power := map[int]int{1: 7, 2: 5, 3: 3, 4: 2, 5: 1}
	if keyExists(4, hand) {
		power[2] = 6
	}
	if keyExists(3, hand) {
		power[3] = 4
	}
	return power[len(hand)]
}

// sort cards accordingly to their power and cards positionts in deck
func sortCards(sl []game, useJoker bool) {
	sort.Slice(sl, func(i, j int) bool {
		if sl[i].power != sl[j].power {
			return sl[i].power < sl[j].power
		}
		var res bool = false
		for k := 0; k < len(sl[i].hand); k++ {
			if sl[i].hand[k] != sl[j].hand[k] {
				return compareCards(sl[i].hand[k], sl[j].hand[k], useJoker)
			}
		}
		return res
	})
}

// compare cards power
func compareCards(r1 byte, r2 byte, useJoker bool) bool {
	var mapPower map[string]int
	if !useJoker {
		mapPower = map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}
	} else {
		mapPower = map[string]int{"A": 1, "K": 2, "Q": 3, "T": 4, "9": 5, "8": 6, "7": 7, "6": 8, "5": 9, "4": 10, "3": 11, "2": 12, "J": 13}
	}
	return mapPower[string(r1)] > mapPower[string(r2)]
}

// find card to which add joker power
func findOptimumCard(hand map[rune]int) rune {
	var maxCard int
	var maxKey rune
	for k, v := range hand {
		if v >= maxCard {
			maxCard = v
			maxKey = k
		}
	}
	return maxKey
}
