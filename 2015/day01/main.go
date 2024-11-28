package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(input []string) int {
	floor := 0
	for i := 0; i < len(input[0]); i++ {
		if input[0][i] == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func part2(input []string) int {
	floor := 0
	for i, r := range input[0] {
		if r == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
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
