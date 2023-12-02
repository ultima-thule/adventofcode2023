package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func cubeConudrum(filename string, calcFun func(int, int, int, int) int) int {
	fmt.Println("=> DataSet: ", filename)

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

func calcSum(game int, green int, red int, blue int) int {
	const maxGreen, maxRed, maxBlue int = 13, 12, 14

	if green <= maxGreen && red <= maxRed && blue <= maxBlue {
		return game
	}
	return 0
}

func calcPower(game int, green int, red int, blue int) int {
	return green * red * blue
}

func main() {
	fmt.Println(cubeConudrum("input_test1.txt", calcSum))

	fmt.Println(cubeConudrum("input1.txt", calcSum))

	fmt.Println(cubeConudrum("input_test2.txt", calcPower))

	fmt.Println(cubeConudrum("input2.txt", calcPower))
}
