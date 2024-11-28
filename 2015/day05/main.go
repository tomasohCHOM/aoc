package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
var invalidStrings = []string{"ab", "cd", "pq", "xy"}

func isStringNice(str string) bool {
	vowelCount, containsDouble := 0, false
	if slices.Contains(vowels, str[0]) {
		vowelCount++
	}
	for i := 1; i < len(str); i++ {
		if slices.Contains(vowels, str[i]) {
			vowelCount++
		}
		if str[i] == str[i-1] {
			containsDouble = true
		}
		if slices.Contains(invalidStrings, string(str[i-1])+string(str[i])) {
			return false
		}
	}
	return vowelCount >= 3 && containsDouble
}

func part1(input []string) int {
	output := 0
	for _, str := range input {
		if isStringNice(str) {
			output += 1
		}
	}
	return output
}

func part2(input []string) int {
	return 0
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
