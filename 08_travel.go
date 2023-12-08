package main

// Day 08 solution
func hauntedWasteland(filename string, calcFun func(map[string][]string, string) int, prepFun func(data []string) map[string][]string) int {
	input, path := readInput08(filename, prepFun)
	return calcFun(input, path)
}

// Solve puzzle no 1
func travel1(input map[string][]string, commands string) int {
	if input == nil || commands == "" {
		return 0
	}

	var res int = 0
	lastNode := "AAA"

	// iterate through commands until ZZZ found
	for {
		found, node, cnt := goSingleRound(input, commands, lastNode)
		res += cnt
		lastNode = node
		if found == true {
			break
		}
	}

	return res
}

func goSingleRound(input map[string][]string, commands string, lastNode string) (bool, string, int) {
	var found bool = false
	var cnt int = 0

	for _, c := range commands {
		command := string(c)
		if command == "L" {
			lastNode = input[lastNode][0]
		}
		if command == "R" {
			lastNode = input[lastNode][1]
		}
		cnt++
		if lastNode == "ZZZ" {
			found = true
			break
		}
	}
	return found, lastNode, cnt
}

// Solve puzzle no 2
func travel2(input map[string][]string, commands string) int {
	if input == nil || commands == "" {
		return 0
	}

	nodes := findStartNodes(input)
	result := calcAllPaths(input, commands, nodes)

	// find LCM among all paths lenghts
	if len(result) > 2 {
		return LCM(result[0], result[1], result...)
	}
	return LCM(result[0], result[len(result)-1])
}

// find all nodes ending with A
func findStartNodes(input map[string][]string) []string {
	res := make([]string, 0)

	for k := range input {
		if k[len(k)-1:] == "A" {
			res = append(res, k)
		}
	}
	return res
}

// calculate single path leghth
func calcAllPaths(input map[string][]string, path string, nodes []string) []int {
	result := make([]int, 0)

	for k := 0; k < len(nodes); k++ {
		i := 0

		endLetter := nodes[k][len(nodes[k])-1:]

		for endLetter != "Z" {
			command := string(path[i%len(path)])
			nodes[k] = input[nodes[k]][getCmdIdx(command)]
			i++

			endLetter = nodes[k][len(nodes[k])-1:]

		}
		result = append(result, i)
	}
	return result
}

func getCmdIdx(letter string) int {
	if letter == "L" {
		return 0
	}
	return 1
}
