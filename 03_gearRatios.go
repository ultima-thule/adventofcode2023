package main

import (
	"bufio"
	"strings"
	"unicode"
)

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

func prepareData(schema []string, text string, isLast bool) []string {
	if len(schema) == 0 {
		schema = append(schema, strings.Repeat(".", len(text)))
	}
	schema = append(schema, text)

	if isLast {
		schema = append(schema, strings.Repeat(".", len(schema[0])))
	}

	if len(schema) > 3 {
		schema = schema[1:]
	}
	return schema
}

func isSymbol(c rune) bool {
	return !unicode.IsDigit(rune(c)) && string(c) != "."
}

func hasSymbols(row string) bool {
	var res bool

	for _, v := range row {
		res = res || isSymbol(v)
	}
	return res
}

func findParts(schema []string) int {
	if len(schema) != 3 {
		return 0
	}

	var result int
	var isAdj bool
	tempNum := ""

	for i, c := range schema[1] {
		isDigit := unicode.IsDigit(c)

		if isDigit {
			tempNum += string(c)
			iLeft := max(0, i-1)
			iRight := min(i+1, len(schema[1])-1)

			for _, v := range schema {
				isAdj = isAdj || hasSymbols(v[iLeft:iRight+1])
			}
		}

		lastItem := len(schema[1]) - 1

		if (isDigit && i == lastItem) || (!isDigit && tempNum != "") {
			if isAdj {
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
	return convert(res)
}

func scanRight(s string) int {
	var res string

	for i := 0; i < len(s) && unicode.IsDigit(rune(s[i])); i++ {
		res = res + string(s[i])
	}
	return convert(res)
}

func getPartNumbers(r string, left int, i int, right int) []int {
	parts := []int{}

	t1 := unicode.IsDigit(rune(r[left]))
	t2 := unicode.IsDigit(rune(r[i]))
	t3 := unicode.IsDigit(rune(r[right]))

	if !t2 {
		if t1 {
			parts = append(parts, scanLeft(r[:left+1]))
		}
		if t3 {
			parts = append(parts, scanRight(r[right:]))
		}
	} else if t1 && t2 && t3 {
		parts = append(parts, scanRight(r[left:]))
	} else if t2 && !t3 {
		parts = append(parts, scanLeft(r[:i+1]))
	} else if t2 && !t1 {
		parts = append(parts, scanRight(r[i:]))
	}

	return parts
}

func findAdjacentGears(engine []string) int {
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
				parts = append(parts, getPartNumbers(v, maxLeft, i, maxRight)...)
			}

			if len(parts) == 2 {
				result += parts[0] * parts[1]
			}
		}
	}
	return result
}
