package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readFile(filename string) *os.File {
	fmt.Println("=> DataSet: ", filename)

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

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func convert(text string) int {
	intVar, err := strconv.Atoi(text)
	if err == nil {
		return intVar
	}
	return 0
}

func convert64(text string) int64 {
	intVar, err := strconv.Atoi(text)
	if err == nil {
		return int64(intVar)
	}
	return 0
}

func readInput(filename string, calcFun func([]string) int, prepFun func(data []string) []string) []string {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	input := []string{}

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	if prepFun != nil {
		input = prepFun(input)
	}

	return input
}

func readInput07(filename string, prepFun func(data []string) []game) []game {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	input := []string{}

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	output := prepFun(input)
	return output
}

func prepData07(input []string) []game {
	if input == nil {
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

// check if key exists
func keyExists(value int, deck map[rune]int) bool {
	for _, v := range deck {
		if v == value {
			return true
		}
	}
	return false
}

func readInput08(filename string, prepFun func(data []string) map[string][]string) (map[string][]string, string) {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	input := []string{}
	var path string = ""

	for fileScanner.Scan() {
		if path == "" {
			path = fileScanner.Text()
			continue
		}
		txt := fileScanner.Text()

		input = append(input, txt)
	}

	output := prepFun(input[1:])
	return output, path
}

func prepData08(input []string) map[string][]string {
	if input == nil {
		return nil
	}

	// fmt.Println(input)

	output := make(map[string][]string, len(input))

	for i := 0; i < len(input); i++ {
		a := strings.Split(input[i], " = ")
		c := strings.ReplaceAll(a[1], "(", "")
		c = strings.ReplaceAll(c, ")", "")
		b := strings.Split(c, ", ")
		output[a[0]] = []string{b[0], b[1]}
	}

	return output
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// split text into int values
func splitToInts(text string) []int {
	tmp := strings.Fields(text)
	values := make([]int, 0, len(tmp))

	for _, raw := range tmp {
		// fmt.Println(raw)
		v, err := strconv.Atoi(raw)
		if err != nil {
			// log.Print(err)
			continue
		}
		values = append(values, v)
	}
	return values
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func prepData09(input string) []int {
	return splitToInts(input)
}

func readInput10(filename string) ([]string, Point) {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	input := []string{}
	pos := Point{x: -1, y: -1}

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		// fmt.Println(txt)
		if pos.y == -1 {
			pos.x++
			var startPos string = `S`
			sent := regexp.MustCompile(startPos)
			ind := sent.FindAllStringIndex(txt, 1)
			if len(ind) > 0 {
				pos.y = ind[0][0]
			}
		}
		fmt.Println(txt)
		input = append(input, txt)
	}

	return input, pos
}

func parseInputIntoBytes(input string) [][]byte {
	splitted := strings.Split(input, "\n")
	grid := make([][]byte, len(splitted))
	for i := range splitted {
		grid[i] = []byte(splitted[i])
	}

	return grid
}

func distance(p1 Point, p2 Point) int {
	return abs(p2.y-p1.y) + abs(p2.x-p1.x)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func abs64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func printTilt(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				fmt.Print("#")
				continue
			}
			if roundRocks[key] {
				fmt.Print("O")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func transposeMatrix(input []string) []string {
	ret := make([]string, len(input[0]))

	for i := 0; i < len(input); i++ {
		t := input[i]
		for j := 0; j < len(t); j++ {
			ret[j] = ret[j] + string(t[j])
		}
	}

	return ret
}

func printVisited(visited map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			keyN := fmt.Sprintf("%d_%d_N", r, c)
			keyS := fmt.Sprintf("%d_%d_S", r, c)
			keyW := fmt.Sprintf("%d_%d_W", r, c)
			keyE := fmt.Sprintf("%d_%d_E", r, c)
			if visited[keyN] || visited[keyS] || visited[keyW] || visited[keyE] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func (d Direction) String() string {
	return []string{"N", "E", "S", "W", "N1", "E1", "S1", "W1", "N2", "E2", "S2", "W2", "N3", "E3", "S3", "W3"}[d]
}

func parseInputIntoNodes(input string) [][]Node {
	splitted := strings.Split(input, "\n")
	grid := make([][]Node, len(splitted))
	for i := range grid {
		grid[i] = make([]Node, len(splitted[0]))
	}
	for i := range splitted {
		for k, v := range splitted[i] {
			conv := convert(string(v))
			grid[i][k] = Node{x: i, y: k, score: conv}
		}
	}

	return grid
}

func parseInput18(input string) []DigPlan {
	splitted := strings.Split(input, "\n")
	grid := make([]DigPlan, 0)

	for _, v := range splitted {
		tmp := strings.Split(v, " ")
		clr := strings.ReplaceAll(tmp[2], "(#", "")
		clr = strings.ReplaceAll(clr, ")", "")
		dp := DigPlan{dir: tmp[0], moves: convert64(string(tmp[1]))}
		grid = append(grid, dp)
	}

	// fmt.Println(grid)

	return grid
}

func parseInput18_2(input string) []DigPlan {
	splitted := strings.Split(input, "\n")
	grid := make([]DigPlan, 0)

	mapDir := map[byte]string{'0': "R", '1': "D", '2': "L", '3': "U"}

	for _, v := range splitted {
		tmp := strings.Split(v, " ")
		clr := strings.ReplaceAll(tmp[2], "(#", "")
		clr = strings.ReplaceAll(clr, ")", "")
		value, _ := strconv.ParseInt(clr[0:5], 16, 64)
		dp := DigPlan{dir: mapDir[clr[5]], moves: value}
		grid = append(grid, dp)
	}

	// fmt.Println(grid)

	return grid
}

func printVisited18(visited map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			keyN := fmt.Sprintf("%d_%d", r, c)
			if visited[keyN] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}
