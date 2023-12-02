package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readFile(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var maxGreen, maxRed, maxBlue int = 13, 12, 14
	var sumGames int

	for fileScanner.Scan() {
		game := matchStr(fileScanner.Text(), "Game (\\d+):")

		green := matchStr(fileScanner.Text(), "(\\d+) green")
		red := matchStr(fileScanner.Text(), "(\\d+) red")
		blue := matchStr(fileScanner.Text(), "(\\d+) blue")

		if green <= maxGreen && red <= maxRed && blue <= maxBlue {
			sumGames += game
		}

		fmt.Println(fileScanner.Text(), " => G: ", green, " R: ", red, "B: ", blue)
	}
	return sumGames
}

func readFile2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var gamesPower, sumGames int

	for fileScanner.Scan() {
		green := matchStr(fileScanner.Text(), "(\\d+) green")
		red := matchStr(fileScanner.Text(), "(\\d+) red")
		blue := matchStr(fileScanner.Text(), "(\\d+) blue")

		gamesPower = green * red * blue

		sumGames += gamesPower

		// fmt.Println(fileScanner.Text(), " => G: ", green, " R: ", red, "B: ", blue)
	}
	return sumGames
}

func matchStr(text string, patt string) int {
	pattern := regexp.MustCompile(patt)
	var greater, numb int

	matches := pattern.FindAllStringSubmatch(text, -1)
	// fmt.Printf("Matches %v", matches)
	// fmt.Println("Text: ", text, " Matches: ", pattern.NumSubexp(), " first: ", matches[0], " - ", matches[1])

	for _, v := range matches {
		// fmt.Printf("%d. %s => %s\n", k, v, v[1])
		numb, _ = strconv.Atoi(v[1])
		greater = max(numb, greater)
	}

	return greater
}

func main() {
	fmt.Println("---- Test DataSet1 ----")
	fmt.Println(readFile("input_test1.txt"))

	fmt.Println("---- DataSet 1 ----")
	fmt.Println(readFile("input1.txt"))

	fmt.Println("---- Test DataSet 2 ----")
	fmt.Println(readFile2("input_test2.txt"))

	fmt.Println("---- DataSet 2 ----")
	fmt.Println(readFile2("input2.txt"))
}
