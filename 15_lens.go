package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	focal int
}

func lens(filename string, puzzle func(string) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		input = append(input, strings.Split(fileScanner.Text(), ",")...)
	}
	for i := 0; i < len(input); i++ {
		res += puzzle(input[i])
	}

	return res
}

func lens_part2(filename string, puzzle func(string) map[string]string) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		input = append(input, strings.Split(fileScanner.Text(), ",")...)
	}

	lenses := make([][]Lens, 256)
	for i := 0; i < 256; i++ {
		lenses = append(lenses, make([]Lens, 0))
	}
	for i := 0; i < len(input); i++ {
		mapRes := puzzle(input[i])
		moveLenses(mapRes, &lenses)
	}

	res = calcFocusingPower(&lenses)

	return res
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
