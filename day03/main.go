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
		// printSlice(window)
	}
	window = append(window, strings.Repeat(".", len(window[0])))[1:]
	result += calcFun(window, row)

	// printSlice(window)

	return result
}

func isSymbol(c byte) bool {
	// fmt.Println(string(c), !unicode.IsDigit(rune(c)), string(c) != ".", !unicode.IsDigit(rune(c)) && string(c) != ".")
	// fmt.Println(!unicode.IsDigit(rune(c)) && string(c) != ".")
	return !unicode.IsDigit(rune(c)) && string(c) != "."
}

func checkAdjacent(engine []string, row int) int {
	var result int

	if len(engine) != 3 {
		return 0
	}

	// printSlice(engine)
	// fmt.Println(engine[1])
	lastNum := ""
	lastAdj := false

	for i, c := range engine[1] {
		// fmt.Printf("%s", string(c))
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
	// fmt.Printf("\n")
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

func checkAdjacentGears(engine []string, row int) int {
	var result int

	if len(engine) != 3 {
		return 0
	}

	// printSlice(engine)
	// fmt.Println(engine[1])

	for i, c := range engine[1] {
		// fmt.Printf("%s", string(c))
		if string(c) == "*" {
			maxLeft := max(0, i-1)
			maxRight := min(i+1, len(engine[1])-1)

			// fmt.Println("==>", string(c), ": L: ", maxLeft, ", R: ", maxRight)
			// fmt.Println("==>", engine[1])

			t1 := unicode.IsDigit(rune(engine[0][maxLeft]))
			t2 := unicode.IsDigit(rune(engine[0][i]))
			t3 := unicode.IsDigit(rune(engine[0][maxRight]))

			// fmt.Println("Check top")
			var parts []int

			if !t2 {
				if t1 {
					parts = append(parts, scanLeft(engine[0][:maxLeft+1]))
				}
				if t3 {
					parts = append(parts, scanRight(engine[0][maxRight:]))
				}
			} else if t1 && t2 && t3 {
				parts = append(parts, scanRight(engine[0][maxLeft:]))
			} else if t2 && !t3 {
				parts = append(parts, scanLeft(engine[0][0:i+1]))
			} else if t2 && !t1 {
				parts = append(parts, scanRight(engine[0][i:len(engine[1])]))
			}

			// fmt.Println("Check mid")
			m1 := unicode.IsDigit(rune(engine[1][maxLeft]))
			m3 := unicode.IsDigit(rune(engine[1][maxRight]))

			if m1 {
				parts = append(parts, scanLeft(engine[1][:maxLeft+1]))
			}
			if m3 {
				parts = append(parts, scanRight(engine[1][maxRight:]))
			}

			// fmt.Println("Check bottom")
			b1 := unicode.IsDigit(rune(engine[2][maxLeft]))
			b2 := unicode.IsDigit(rune(engine[2][i]))
			b3 := unicode.IsDigit(rune(engine[2][maxRight]))

			if !b2 {
				if b1 {
					parts = append(parts, scanLeft(engine[2][:maxLeft+1]))
				}
				if b3 {
					parts = append(parts, scanRight(engine[2][maxRight:]))
				}
			} else if b1 && b2 && b3 {
				parts = append(parts, scanRight(engine[2][maxLeft:]))
			} else if b2 && !b3 {
				parts = append(parts, scanLeft(engine[2][:i+1]))
			} else if b2 && !b1 {
				parts = append(parts, scanRight(engine[2][i:]))
			}

			// fmt.Println("Parts in row ", row, "len ", len(parts), parts)

			if len(parts) == 2 {
				result += parts[0] * parts[1]
				// fmt.Println(row, ": ", parts[0], parts[1])
			}

			// fmt.Printf("%s:\n%s %s %s\n%t * %t\n%t %t %t\n", string(c), t1, t2, t3, m1, m3, b1, b2, b3)
			// fmt.Printf("%s:\n%t %t %t\n%t * %t\n%t %t %t\n", string(c), t1, t2, t3, m1, m3, b1, b2, b3)

			// if t1 || t2 || t3 || m1 || m3 || b1 || b2 || b3 {
			// 	lastAdj = true
			// }
			// fmt.Printf("%s: %t\n", string(c), lastAdj)

			//} else if (!unicode.IsDigit(c) || i == len(engine[1])-1) && lastNum != "" {
			// if i == len(engine[1])-1 {
			// 	if lastAdj == true {
			// 		// fmt.Println("entered!", string(c))
			// 		intVar, err := strconv.Atoi(lastNum)
			// 		if err == nil {
			// 			result += intVar
			// 		}
			// 		// fmt.Println("char ", string(c), " end of number ", lastNum, " adjacent ", lastAdj, " converted ", intVar, " result ", result)
			// 		// fmt.Printf("%d ", intVar)
			// 	}
			// 	lastNum = ""
			// 	lastAdj = false
			// }
		}
	}
	// fmt.Println("Result:", result)
	return result
}

func main() {
	fmt.Println(gearRatios("input_test1.txt", checkAdjacent))

	fmt.Println(gearRatios("input1.txt", checkAdjacent))

	fmt.Println(gearRatios("input_test3.txt", checkAdjacentGears))

	fmt.Println(gearRatios("input2.txt", checkAdjacentGears))
}
