package main

import (
	"flag"
	"fmt"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "03", "day of AoC competition in format dd with trailing zero")
	flag.Parse()

	switch day {
	case "01":
		fmt.Println(trebuchet("day"+day+"/input_test1.txt", 1))
		fmt.Println(trebuchet("day"+day+"/input1.txt", 1))
		fmt.Println(trebuchet("day"+day+"/input_test2.txt", 2))
		fmt.Println(trebuchet("day"+day+"/input2.txt", 2))
	case "02":
		fmt.Println(cubeConudrum("day"+day+"/input_test1.txt", calcSum))
		fmt.Println(cubeConudrum("day"+day+"/input1.txt", calcSum))
		fmt.Println(cubeConudrum("day"+day+"/input_test2.txt", calcPower))
		fmt.Println(cubeConudrum("day"+day+"/input2.txt", calcPower))
	case "03":
		fmt.Println(gearRatios("day"+day+"/input_test1.txt", findParts))
		fmt.Println(gearRatios("day"+day+"/input1.txt", findParts))
		fmt.Println(gearRatios("day"+day+"/input_test2.txt", findAdjacentGears))
		fmt.Println(gearRatios("day"+day+"/input2.txt", findAdjacentGears))
	}
}
