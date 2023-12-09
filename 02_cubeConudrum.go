package main

import (
	"regexp"
	"strconv"
)

func puzzle02_1(game int, green int, red int, blue int) int {
	const maxGreen, maxRed, maxBlue int = 13, 12, 14

	if green <= maxGreen && red <= maxRed && blue <= maxBlue {
		return game
	}
	return 0
}

func puzzle02_2(game int, green int, red int, blue int) int {
	return green * red * blue
}

func findMax(text string, patt string) int {
	pattern := regexp.MustCompile(patt)
	var greater int

	matches := pattern.FindAllStringSubmatch(text, -1)

	for _, v := range matches {
		numb, err := strconv.Atoi(v[1])
		if err == nil {
			greater = max(numb, greater)
		}
	}

	return greater
}

func extractData(text string) (int, int, int, int) {
	game := findMax(text, "Game (\\d+):")

	green := findMax(text, "(\\d+) green")
	red := findMax(text, "(\\d+) red")
	blue := findMax(text, "(\\d+) blue")

	return game, green, red, blue
}
