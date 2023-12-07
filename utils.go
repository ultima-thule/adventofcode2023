package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func readInput07(filename string, calcFun func([]game) int, prepFun func(data []string) []game) []game {
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
