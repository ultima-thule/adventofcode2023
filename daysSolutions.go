package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Day 18 solution
func lavaduct(filename string, calcFun func([]DigPlan) int64) int64 {
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInput18(content)

	res := calcFun(grid)

	defer timeTrack(time.Now(), "lavaduct")
	fmt.Println()

	return res
}

// Day 18 solution
func lavaduct_2(filename string, calcFun func([]DigPlan) int64) int64 {
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInput18_2(content)

	res := calcFun(grid)

	defer timeTrack(time.Now(), "lavaduct")
	fmt.Println()

	return res
}

// Day 17 solution
func crucible(filename string, calcFun func(*[][]Node) int) int {
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInputIntoNodes(content)

	res := calcFun(&grid)

	defer timeTrack(time.Now(), "crucible")
	fmt.Println()

	return res
}

// Day 16 solution
func beam(filename string, calcFun func([][]byte) int) int {
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInputIntoBytes(content)

	res := calcFun(grid)

	defer timeTrack(time.Now(), "beam")
	fmt.Println()

	return res
}

// solution day 13
func points(filename string, calcData func([]int64, []int64) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		t := fileScanner.Text()
		if t == "" {
			dataRow, dataCol := prepData13(input)
			res += calcData(dataRow, dataCol)
			input = make([]string, 0)
		} else {
			input = append(input, t)
		}
	}
	dataRow, dataCol := prepData13(input)
	res += calcData(dataRow, dataCol)

	return res
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

func lens_part2(filename string) int {
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
		mapRes := prepareData15(input[i])
		puzzle15_2(mapRes, &lenses)
	}

	res = calcFocusingPower(&lenses)

	return res
}

// Day 11 solution
func cosmic(filename string, calcFun func([][]byte, int) int, expRate int) int {
	// Read the contents of the file
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInputIntoBytes(content)

	res := calcFun(grid, expRate)

	defer timeTrack(time.Now(), "cosmic")
	fmt.Println()

	return res
}

// Day 10 solution
func pipeMaze(filename string, calcFun func([]string, Point) int) int {
	input, pos := readInput10(filename)
	return calcFun(input, pos)
}

// Day 09 solution
func mirage(filename string, calcFun func([]int) int, prepFun func(string) []int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int

	for fileScanner.Scan() {
		text := fileScanner.Text()
		result += calcFun(prepFun(text))
	}
	defer timeTrack(time.Now(), "mirage")

	return result
}

// Day 08 solution
func hauntedWasteland(filename string, calcFun func(map[string][]string, string) int, prepFun func(data []string) map[string][]string) int {
	input, path := readInput08(filename, prepFun)
	return calcFun(input, path)
}

// Day 07 solution
func camelCards(filename string, calcFun func([]game, bool) int, prepFun func(data []string) []game, useJoker bool) int {
	input := readInput07(filename, prepFun)
	return calcFun(input, useJoker)
}

// Day 06 solution
func waitForIt(filename string, calcFun func([]string) int, prepFun func(data []string) []string) int {
	input := readInput(filename, calcFun, prepFun)

	return calcFun(input)
}

// Day 05 solution
func fertilizer(filename string, calcFun func([]string) int, prepFun func(data []string) []string) int {
	input := readInput(filename, calcFun, prepFun)

	return calcFun(input)
}

// Day 04 solution
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

// Day 03 solution
func gearRatios(filename string, calcFun func([]string) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int
	var schema []string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		schema = prepareData(schema, text, false)
		result += calcFun(schema)
	}
	// last row
	schema = prepareData(schema, "", true)
	result += calcFun(schema)

	return result
}

// Day 02 solution
func cubeConudrum(filename string, calcFun func(int, int, int, int) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var result int

	for fileScanner.Scan() {
		game, green, red, blue := extractData(fileScanner.Text())
		result += calcFun(game, green, red, blue)
	}
	return result
}

// Day 01 solution
func trebuchet(filename string, version int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var first, last, result int

	for fileScanner.Scan() {
		if version == 1 {
			first, last = parseDigitsOnly(fileScanner.Text())
		} else {
			first, last = parseDigitsMixed(fileScanner.Text())
		}
		result += first*10 + last
	}

	return result
}
