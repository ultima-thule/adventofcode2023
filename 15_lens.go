package main

import (
	"bufio"
	"strings"
)

func lens(filename string) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		input = append(input, strings.Split(fileScanner.Text(), ",")...)
	}
	for i := 0; i < len(input); i++ {
		res += puzzle15(input[i])
	}

	return res
}

// Solve puzzle no 14
func puzzle15(input string) int {
	res := 0

	for i := 0; i < len(input); i++ {
		r := rune(input[i])
		res = ((res + int(r)) * 17 % 256)

	}
	return res
}
