package main

import (
	"fmt"
	"os"
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
