package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if input == nil || len(input) < 2 {
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
