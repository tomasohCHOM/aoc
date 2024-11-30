package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
var invalidStrings = []string{"ab", "cd", "pq", "xy"}

func isStringNice1(str string) bool {
	vowelCount, containsDouble := 0, false
	for i := 0; i < len(str); i++ {
		if slices.Contains(vowels, str[i]) {
			vowelCount++
		}
		if i != 0 && str[i] == str[i-1] {
			containsDouble = true
		}
		if i != 0 && slices.Contains(invalidStrings, string(str[i-1])+string(str[i])) {
			return false
		}
	}
	return vowelCount >= 3 && containsDouble
}

func isStringNice2(str string) bool {
	hasDuplicatePair := false
	hasSurrounding := false
	pairs := map[string]int{}

	for i := 0; i < len(str); i++ {
		if i > 1 && str[i-2] == str[i] {
			hasSurrounding = true
		}
		if i > 0 {
			pair := string(str[i-1]) + string(str[i])
			index, ok := pairs[pair]
			if !ok {
				pairs[pair] = i
			} else if i > index+1 {
				hasDuplicatePair = true
			}
		}
	}
	return hasDuplicatePair && hasSurrounding
}

func part1(input []string) int {
	output := 0
	for _, str := range input {
		if isStringNice1(str) {
			output += 1
		}
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, str := range input {
		if isStringNice2(str) {
			output += 1
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
