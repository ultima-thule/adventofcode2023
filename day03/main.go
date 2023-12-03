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

func gearRatios(filename string) int {
	fmt.Println("=> DataSet: ", filename)

	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var result int

	var window []string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(window) == 0 {
			window = append(window, strings.Repeat(".", len(text)))
		}
		window = append(window, text)

		if len(window) > 3 {
			window = window[1:]
		}
		result += checkAdjacent(window)
		// printSlice(window)
	}
	window = append(window, strings.Repeat(".", len(window[0])))[1:]
	result += checkAdjacent(window)

	// printSlice(window)

	return result
}

func isSymbol(c byte) bool {
	// fmt.Println(string(c), !unicode.IsDigit(rune(c)), string(c) != ".", !unicode.IsDigit(rune(c)) && string(c) != ".")
	// fmt.Println(!unicode.IsDigit(rune(c)) && string(c) != ".")
	return !unicode.IsDigit(rune(c)) && string(c) != "."
}

func checkAdjacent(engine []string) int {
	var result int

	if len(engine) != 3 {
		return 0
	}

	// printSlice(engine)
	// fmt.Println(engine[1])
	lastNum := ""
	lastAdj := false

	for i, c := range engine[1] {
		fmt.Printf("%s", string(c))
		if unicode.IsDigit(c) {
			lastNum += string(c)
			maxLeft := max(0, i-1)
			maxRight := min(i+1, len(engine[1])-1)

			// fmt.Println(string(c), ": L: ", maxLeft, ", R: ", maxRight)

			t1 := isSymbol(engine[0][maxLeft])
			t2 := isSymbol(engine[0][i])
			t3 := isSymbol(engine[0][maxRight])

			m1 := isSymbol(engine[1][maxLeft])
			m3 := isSymbol(engine[1][maxRight])

			b1 := isSymbol(engine[2][maxLeft])
			b2 := isSymbol(engine[2][i])
			b3 := isSymbol(engine[2][maxRight])
			// fmt.Printf("%s:\n%s %s %s\n%t * %t\n%t %t %t\n", string(c), t1, t2, t3, m1, m3, b1, b2, b3)
			// fmt.Printf("%s:\n%t %t %t\n%t * %t\n%t %t %t\n", string(c), t1, t2, t3, m1, m3, b1, b2, b3)

			if t1 || t2 || t3 || m1 || m3 || b1 || b2 || b3 {
				lastAdj = true
			}
			// fmt.Printf("%s: %t\n", string(c), lastAdj)

			//} else if (!unicode.IsDigit(c) || i == len(engine[1])-1) && lastNum != "" {
			if i == len(engine[1])-1 {
				if lastAdj == true {
					// fmt.Println("entered!", string(c))
					intVar, err := strconv.Atoi(lastNum)
					if err == nil {
						result += intVar
					}
					// fmt.Println("char ", string(c), " end of number ", lastNum, " adjacent ", lastAdj, " converted ", intVar, " result ", result)
					// fmt.Printf("%d ", intVar)
				}
				lastNum = ""
				lastAdj = false
			}
		} else if !unicode.IsDigit(c) && lastNum != "" {
			// fmt.Println("entered!", string(c))
			if lastAdj == true {
				// fmt.Println("char ", string(c), " end of number ", lastNum, " adjacent ", lastAdj)
				intVar, err := strconv.Atoi(lastNum)
				if err == nil {
					result += intVar
				}
				// fmt.Println("char ", string(c), " end of number ", lastNum, " adjacent ", lastAdj, " converted ", intVar, " result ", result)
				// fmt.Printf("%d ", intVar)
			}
			lastNum = ""
			lastAdj = false
		}
	}
	fmt.Printf("\n")
	return result
}

func main() {
	fmt.Println(gearRatios("input_test1.txt"))

	fmt.Println(gearRatios("input1.txt"))

	// fmt.Println(gearRatios("input_test2.txt", gearRatios))

	// fmt.Println(gearRatios("input2.txt", gearRatios))
}
