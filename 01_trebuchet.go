package main

import (
	"strings"
)

func parseDigitsOnly(text string) (int, int) {
	var first, last int

	for _, letter := range text {
		if letter >= '0' && letter <= '9' {
			last = int(letter - '0')

			if first == 0 {
				first = last
			}
		}
	}

	return first, last
}

func parseDigitsMixed(text string) (int, int) {
	var first, last int = 0, 0

	digitsMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

LOOP1:
	// find first digit
	for i, letter := range text {
		if letter >= '0' && letter <= '9' {
			first = int(letter - '0')
			break
		}

		for k, v := range digitsMap {
			if strings.HasPrefix(text[i:], k) {
				first = v
				break LOOP1
			}
		}
	}

	// find last digit
LOOP2:
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] >= '0' && text[i] <= '9' {
			last = int(text[i] - '0')
			break
		}
		for k, v := range digitsMap {
			if strings.HasSuffix(text[:i+1], k) {
				last = v
				break LOOP2
			}
		}
	}

	return first, last
}
