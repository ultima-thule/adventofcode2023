package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func points(filename string) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		t := fileScanner.Text()
		if t == "" {
			dataRow, dataCol := prepData13(input)
			res += puzzle13(dataRow, dataCol)
			input = make([]string, 0)
		} else {
			input = append(input, t)
		}
	}
	dataRow, dataCol := prepData13(input)
	res += puzzle13(dataRow, dataCol)

	return res
}

func points2(filename string) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		t := fileScanner.Text()
		// fmt.Println(t)
		if t == "" {
			fmt.Println("\nMap found")
			oldRow := make(map[int]bool, 0)
			oldCol := make(map[int]bool, 0)
			dataRow, dataCol := prepData13(input)
			puzzle13_2(dataRow, dataCol, oldRow, oldCol)
			fmt.Println("Original reflections: ", oldRow, oldCol)

			for i := 0; i < len(input); i++ {
				// fmt.Println("Długość:", len(input[0]))
				for j := 0; j < len(input[0]); j++ {
					fmt.Println("1: Point: [", i, j, "]")
					dataRow, dataCol := prepData13_2(input, i, j)
					res += puzzle13_2(dataRow, dataCol, oldRow, oldCol)
				}
			}
			input = make([]string, 0)
		} else {
			input = append(input, t)
		}
	}

	fmt.Println("\n Map found")
	oldRow := make(map[int]bool, 0)
	oldCol := make(map[int]bool, 0)
	dataRow, dataCol := prepData13(input)
	puzzle13_2(dataRow, dataCol, oldRow, oldCol)
	fmt.Println("Original reflections: ", oldRow, oldCol)

	for i := 0; i < len(input); i++ {
		// fmt.Println("Długość:", len(input[0]))
		for j := 0; j < len(input[0]); j++ {
			// fmt.Println("1: Point: [", i, j, "]")
			dataRow, dataCol := prepData13_2(input, i, j)
			res += puzzle13_2(dataRow, dataCol, oldRow, oldCol)
		}
	}

	return res
}

func transposeMirror(input []string) []string {
	ret := make([]string, len(input[0]))

	for i := 0; i < len(input); i++ {
		t := input[i]
		for j := 0; j < len(t); j++ {
			ret[j] = ret[j] + string(t[j])
		}
	}

	return ret
}

func convertToInt(input []string) []int64 {
	// fmt.Println("Puzzle conversion")
	ret := make([]int64, len(input))

	for k, v := range input {
		tmp := strings.ReplaceAll(v, ".", "0")
		tmp = strings.ReplaceAll(tmp, "#", "1")
		// fmt.Println(tmp)
		i, _ := strconv.ParseInt(tmp, 2, 64)
		ret[k] = i
	}
	// fmt.Println()

	return ret
}

func convertToInt_2(input []string, i int, j int, isTransposed bool) []int64 {
	// fmt.Println("Puzzle conversion", input, i, j)
	ret := make([]int64, len(input))

	for k, v := range input {
		text := v
		if !isTransposed {
			if k == i {
				// fmt.Println("Converting ", v, ", at index ", i, j)
				// fmt.Println("Converting row", text[:j])
				// fmt.Println("Converting row", repl(text[j]))
				// fmt.Println("Converting idx", min(j+1, len(input[i])-1))
				text = text[:j] + repl(text[j]) + text[min(j+1, len(v)-1):]
				// fmt.Println("Converted ", text)
			}
		} else {
			if k == j {
				// fmt.Println("Converting ", v, ", at index ", j, i)
				// fmt.Println("Converting row", text[:j])
				// fmt.Println("Converting row", repl(text[j]))
				// fmt.Println("Converting idx", min(j+1, len(input[i])-1))
				text = text[:i] + repl(text[i]) + text[min(i+1, len(v)-1):]
				// fmt.Println("Converted ", text)
			}
		}
		tmp := strings.ReplaceAll(text, ".", "0")
		tmp = strings.ReplaceAll(tmp, "#", "1")
		fmt.Println(tmp)
		num, _ := strconv.ParseInt(tmp, 2, 64)
		ret[k] = num
	}
	fmt.Println()

	return ret
}

func repl(r byte) string {
	if string(r) == "." {
		return "#"
	}
	return "."
}

func prepData13(input []string) ([]int64, []int64) {
	retRow := convertToInt(input)

	cols := transposeMirror(input)
	retCol := convertToInt(cols)

	return retRow, retCol
}

func prepData13_2(input []string, i int, j int) ([]int64, []int64) {
	retRow := convertToInt_2(input, i, j, false)

	cols := transposeMirror(input)
	retCol := convertToInt_2(cols, i, j, true)

	return retRow, retCol
}

// Solve puzzle no 1
func puzzle13(inputRow []int64, inputCol []int64) int {
	res := 0
	// fmt.Println("Puzzle by rows")
	// fmt.Println(inputRow)
	// fmt.Println()

	// fmt.Println("Puzzle by cols")
	// fmt.Println(inputCol)
	// fmt.Println()

	res = findMirror(inputRow, inputCol)

	return res
}

func puzzle13_2(inputRow []int64, inputCol []int64, oldRow map[int]bool, oldCol map[int]bool) int {
	res := 0
	// fmt.Println("Puzzle by rows")
	// fmt.Println(inputRow)
	// fmt.Println()

	// fmt.Println("Puzzle by cols")
	// fmt.Println(inputCol)
	// fmt.Println()

	res = findMirror2(inputRow, inputCol, oldRow, oldCol)

	return res
}

func findMirror(inputRow []int64, inputCol []int64) int {
	res := 0
	reflRow := findReflectionLines(inputRow)
	// fmt.Println("Reflections:", reflRow)
	if len(reflRow) > 0 {
		for i := 0; i < len(reflRow); i++ {
			fmt.Println("Reflection in row: ", reflRow[i])

			bRow := checkReflectionLine(reflRow[i], inputRow)
			if bRow {
				res += 100 * (reflRow[i] + 1)
			}
		}
	} else {
		fmt.Println("No row reflection")
	}

	reflColumn := findReflectionLines(inputCol)
	if len(reflColumn) > 0 {
		for i := 0; i < len(reflColumn); i++ {
			fmt.Println("Reflection in col: ", reflColumn[i])
			bCol := checkReflectionLine(reflColumn[i], inputCol)
			if bCol {
				res += (reflColumn[i] + 1)
			}
		}
	} else {
		fmt.Println("No col reflection")
	}

	return res
}

func findMirror2(inputRow []int64, inputCol []int64, oldRow map[int]bool, oldCol map[int]bool) int {
	res := 0
	reflRow := findReflectionLines(inputRow)
	// fmt.Println("Reflections:", reflRow)
	if len(reflRow) > 0 {
		for i := 0; i < len(reflRow); i++ {
			if !oldRow[reflRow[i]] {
				bRow := checkReflectionLine(reflRow[i], inputRow)
				if bRow {
					fmt.Println("Reflection found in row: ", reflRow[i], "res=", res, "adding ", 100*(reflRow[i]+1))
					res += 100 * (reflRow[i] + 1)
					oldRow[reflRow[i]] = true
				}
			}
		}
	} else {
		// fmt.Println("No row reflection")
	}

	reflColumn := findReflectionLines(inputCol)
	if len(reflColumn) > 0 {
		for i := 0; i < len(reflColumn); i++ {
			if !oldCol[reflColumn[i]] {

				bCol := checkReflectionLine(reflColumn[i], inputCol)
				if bCol {
					fmt.Println("Reflection found in col: ", reflColumn[i], "res=", res, " adding ", (reflColumn[i] + 1))
					res += (reflColumn[i] + 1)
					oldCol[reflColumn[i]] = true
				}
			}
		}
	} else {
		// fmt.Println("No col reflection")
	}

	return res
}

func findReflectionLines(input []int64) []int {
	res := make([]int, 0)

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			res = append(res, i)
		}
	}

	return res
}

func checkReflectionLine(reflRow int, input []int64) bool {
	ret := true

	leftLen := reflRow
	rigthLen := len(input) - 1 - reflRow - 1

	// fmt.Println("Data:", reflRow, len(input), leftLen, rigthLen)

	cnt := min(leftLen, rigthLen)

	for i := 0; i <= cnt; i++ {
		// fmt.Println("Comparing", input[reflRow-i], input[reflRow+1+i])
		ret = ret && input[reflRow-i] == input[reflRow+1+i]
	}

	// fmt.Println("Result of check: ", ret)

	return ret
}
