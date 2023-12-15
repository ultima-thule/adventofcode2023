package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Lens struct {
	label string
	focal int
}

// Solve puzzle no 15 part 1
func puzzle15(input string) int {
	res := 0

	for i := 0; i < len(input); i++ {
		r := rune(input[i])
		res = ((res + int(r)) * 17 % 256)

	}
	return res
}

// Solve puzzle no 15 part 2
func puzzle15_2(input string) map[string]string {
	paramsMap := decodeLabel(input)
	paramsMap["box"] = fmt.Sprint(puzzle15(paramsMap["lbl"]))
	return paramsMap
}

// move lenses according to command
func moveLenses(paramsMap map[string]string, lenses *[][]Lens) {
	box, _ := strconv.Atoi(paramsMap["box"])
	cmd := paramsMap["cmd"]
	foc := paramsMap["focal"]
	lab := paramsMap["lbl"]

	switch cmd {
	case "=":
		focal := convert((foc))
		newLens := Lens{label: lab, focal: focal}

		idx := lensIndex(lenses, box, lab)
		if idx == -1 {
			(*lenses)[box] = append((*lenses)[box], newLens)
		} else {
			(*lenses)[box][idx].focal = focal
		}
	case "-":
		idx := lensIndex(lenses, box, lab)
		if idx > -1 {
			(*lenses)[box] = append((*lenses)[box][0:idx], (*lenses)[box][idx+1:]...)
		}
	}
}

// search for index of lens with given label
func lensIndex(lenses *[][]Lens, box int, lab string) int {
	idx := -1
	for i := 0; i < len((*lenses)[box]); i++ {
		if (*lenses)[box][i].label == lab {
			idx = i
			break
		}
	}
	return idx
}

// calculate focusing power: box+1 * lens index * lens focal
func calcFocusingPower(lenses *[][]Lens) int {
	res := 0
	for b := 0; b < 256; b++ {
		for i := 0; i < len((*lenses)[b]); i++ {
			power := (1 + b) * (i + 1) * ((*lenses)[b][i].focal)
			res += power
		}
	}
	return res
}

// parse sequence of operations
func decodeLabel(input string) map[string]string {
	compRegEx := regexp.MustCompile("(?P<lbl>[a-z]+)(?P<cmd>[-=])(?P<focal>\\d*)")

	match := compRegEx.FindStringSubmatch(input)

	paramsMap := make(map[string]string, 0)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
