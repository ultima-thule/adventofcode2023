package main

import (
	"fmt"
)

func main() {
	fmt.Println(gearRatios("day03/input_test1.txt", findParts))

	fmt.Println(gearRatios("day03/input1.txt", findParts))

	fmt.Println(gearRatios("day03/input_test2.txt", findAdjacentGears))

	fmt.Println(gearRatios("day03/input2.txt", findAdjacentGears))
}
