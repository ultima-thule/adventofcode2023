package main

// Day 08 solution
func hauntedWasteland(filename string, calcFun func(map[string][]string, string) int, prepFun func(data []string) map[string][]string) int {
	input, path := readInput08(filename, prepFun)
	return calcFun(input, path)
}

// Solve puzzle no 1
func travel(input map[string][]string, path string) int {
	if input == nil || path == "" {
		return 0
	}

	var res int = 0
	lastNode := "AAA"
	// fmt.Println("Path:", path)

	for {
		found, node, cnt := goThroughPath(input, path, lastNode)
		// fmt.Println("Found: ", found, " node: ", node, " cnt ", cnt, " res ", res)
		res += cnt
		lastNode = node
		if found == true {
			break
		}
	}

	return res
}

// Solve puzzle no 2
func travel2(input map[string][]string, path string) int {
	if input == nil || path == "" {
		return 0
	}

	var res int = 0
	nodes := findNodes(input)
	// fmt.Println("Starting nodes:", nodes)

	res = goThroughPath2(input, path, nodes)

	return res
}

func findNodes(input map[string][]string) []string {
	res := make([]string, 0)

	for k := range input {
		if k[len(k)-1:] == "A" {
			res = append(res, k)
		}
	}

	return res
}

func goThroughPath(input map[string][]string, path string, lastNode string) (bool, string, int) {
	var found bool = false
	var cnt int = 0

	for _, c := range path {
		// fmt.Println("Let:", string(c))
		if string(c) == "L" {
			lastNode = input[lastNode][0]
		}
		if string(c) == "R" {
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

func goThroughPath2(input map[string][]string, path string, nodes []string) int {
	// var found bool = false
	// var cnt int = 0

	// fmt.Println("Nodes list: ", nodes)
	result := make([]int, 0)

	for k := 0; k < len(nodes); k++ {
		i := 0

		letter := nodes[k][len(nodes[k])-1:]
		// fmt.Println("[", i, "] Last letter: ", letter)

		for letter != "Z" {
			item := string(path[i%len(path)])
			// fmt.Println("Item: ", item)

			if item == "L" {
				// fmt.Println("Go left to: ", input[nodes[k]][0])
				nodes[k] = input[nodes[k]][0]
			}
			if item == "R" {
				// fmt.Println("Go right to: ", input[nodes[k]][1])
				nodes[k] = input[nodes[k]][1]
			}
			i++

			letter = nodes[k][len(nodes[k])-1:]
			// fmt.Println("Last letter: ", letter)

		}
		result = append(result, i)
		// fmt.Println(result)
	}

	// return LCM(result[0], result[1], result[2], result[3], result[4], result[5])
	if len(result) > 2 {
		return LCM(result[0], result[1], result...)
	}
	return LCM(result[0], result[len(result)-1])
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
