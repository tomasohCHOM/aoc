package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coords struct {
	r, c int
}

func parseMap(original []string) [][]int {
	parsedMap := [][]int{}
	for r := range original {
		parsedMap = append(parsedMap, []int{})
		for c := range original[r] {
			val, _ := strconv.Atoi(string(original[r][c]))
			parsedMap[r] = append(parsedMap[r], val)
		}
	}
	return parsedMap
}

func part1(input []string) int {
	parsedMap := parseMap(input)
	output := 0

	var dfs func(i, r, c int, seen map[coords]bool, input [][]int)
	dfs = func(i, r, c int, seen map[coords]bool, input [][]int) {
		if min(r, c) < 0 || r >= len(input) || c >= len(input[0]) || seen[coords{r, c}] || input[r][c] != i {
			return
		}
		if i == 9 {
			output += 1
		}
		seen[coords{r, c}] = true
		for _, newCoords := range []coords{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
			dfs(i+1, newCoords.r, newCoords.c, seen, input)
		}
		return
	}
	for r := range parsedMap {
		for c := range parsedMap[r] {
			dfs(0, r, c, map[coords]bool{}, parsedMap)
		}
	}
	return output
}

func part2(input []string) int {
	parsedMap := parseMap(input)
	output := 0
	var dfs func(i, r, c int, seen map[coords]bool, input [][]int)
	dfs = func(i, r, c int, seen map[coords]bool, input [][]int) {
		if min(r, c) < 0 || r >= len(input) || c >= len(input[0]) || seen[coords{r, c}] || input[r][c] != i {
			return
		}
		if i == 9 {
			output += 1
		}
		seen[coords{r, c}] = true
		for _, newCoords := range []coords{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
			dfs(i+1, newCoords.r, newCoords.c, seen, input)
		}
		seen[coords{r, c}] = false
	}
	for r := range parsedMap {
		for c := range parsedMap[r] {
			dfs(0, r, c, map[coords]bool{}, parsedMap)
		}
	}
	return output
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
