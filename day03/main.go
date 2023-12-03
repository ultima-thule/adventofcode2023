package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func gearRatios(filename string, calcFun func([]string, int) int) int {
	fmt.Println("=> DataSet: ", filename)

	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var result int

	var window []string

	var row int
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(window) == 0 {
			window = append(window, strings.Repeat(".", len(text)))
		}
		window = append(window, text)

		if len(window) > 3 {
			window = window[1:]
		}
		result += calcFun(window, row)
		row++
	}
	window = append(window, strings.Repeat(".", len(window[0])))[1:]
	result += calcFun(window, row)

	return result
}

func isSymbol(c rune) bool {
	return !unicode.IsDigit(rune(c)) && string(c) != "."
}

func convert(text string) int {
	intVar, err := strconv.Atoi(text)
	if err == nil {
		return intVar
	}
	return 0
}

func hasSymbols(row string) bool {
	var res bool

	for _, v := range row {
		res = res || isSymbol(v)
	}

	return res
}

func checkPartNums(engine []string, row int) int {
	if len(engine) != 3 {
		return 0
	}

	var result int
	var isAdj bool
	tempNum := ""

	for i, c := range engine[1] {
		if unicode.IsDigit(c) {
			tempNum += string(c)
			iLeft := max(0, i-1)
			iRight := min(i+1, len(engine[1])-1)

			for _, v := range engine {
				isAdj = isAdj || hasSymbols(v[iLeft:iRight+1])
			}
		}
		if (unicode.IsDigit(c) && i == len(engine[1])-1) || (!unicode.IsDigit(c) && tempNum != "") {
			if isAdj == true {
				result += convert(tempNum)
			}
			tempNum = ""
			isAdj = false
		}
	}
	return result
}

func scanLeft(s string) int {
	var res string
	for i := len(s) - 1; i >= 0 && unicode.IsDigit(rune(s[i])); i-- {
		res = string(s[i]) + res
	}
	intVar, err := strconv.Atoi(res)
	if err == nil {
		// fmt.Println("Scan right: ", intVar)
		return intVar
	}
	return 0
}

func scanRight(s string) int {
	var res string

	for i := 0; i < len(s) && unicode.IsDigit(rune(s[i])); i++ {
		res += string(s[i])
	}
	intVar, err := strconv.Atoi(res)
	if err == nil {
		// fmt.Println("Scan right: ", intVar)
		return intVar
	}
	return 0
}

func getNumbers(r string, maxLeft int, i int, maxRight int) []int {
	parts := []int{}

	t1 := unicode.IsDigit(rune(r[maxLeft]))
	t2 := unicode.IsDigit(rune(r[i]))
	t3 := unicode.IsDigit(rune(r[maxRight]))

	if !t2 {
		if t1 {
			parts = append(parts, scanLeft(r[:maxLeft+1]))
		}
		if t3 {
			parts = append(parts, scanRight(r[maxRight:]))
		}
	} else if t1 && t2 && t3 {
		parts = append(parts, scanRight(r[maxLeft:]))
	} else if t2 && !t3 {
		parts = append(parts, scanLeft(r[0:i+1]))
	} else if t2 && !t1 {
		parts = append(parts, scanRight(r[i:]))
	}

	return parts
}

func checkAdjacentGears(engine []string, row int) int {
	var result int

	if len(engine) != 3 {
		return 0
	}

	for i, c := range engine[1] {
		if string(c) == "*" {
			maxLeft := max(0, i-1)
			maxRight := min(i+1, len(engine[1])-1)

			parts := []int{}

			for _, v := range engine {
				parts = append(parts, getNumbers(v, maxLeft, i, maxRight)...)
			}

			if len(parts) == 2 {
				result += parts[0] * parts[1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(gearRatios("input_test1.txt", checkPartNums))

	fmt.Println(gearRatios("input1.txt", checkPartNums))

	fmt.Println(gearRatios("input_test2.txt", checkAdjacentGears))

	fmt.Println(gearRatios("input2.txt", checkAdjacentGears))
}
